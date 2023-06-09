package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

const (
	batchSize = 10
)

var (
	_logrus       *logrus.Logger
	elasticClient *elastic.Client
	logsCh        chan *LogData
	wg            sync.WaitGroup
)

type LogData struct {
	Locate         string `json:"locate"`
	ServiceName    string `json:"service_name"`
	LogID          string `json:"log_id"`
	LogLevel       string `json:"log_level"`
	LogMessage     string `json:"log_message"`
	ExecutionStack string `json:"execution_stack"`
	FileLine       string `json:"file_line"`
	Timestamp      string `json:"timestamp"`
}

func NewLogger(elasticURL []string, file string) error {
	// Initialize logrus
	_logrus = logrus.New()
	_logrus.SetFormatter(&logrus.JSONFormatter{})

	// Initialize elastic client

	var err error
	// 初始化elasticClient
	elasticClient, err = elastic.NewClient(
		elastic.SetURL("https://es-ca7nr12x.public.tencentelasticsearch.com:9200"),
		elastic.SetSniff(false), // SetSniff需要被设置为false，因为你正在使用公网IP
		elastic.SetBasicAuth("elastic", "FOWrYfQbfnRa1_WMepPk"),
		elastic.SetScheme("https"), // 设置使用https
	)
	if err != nil {
		return err
	}

	// Set output file
	// Open the log file
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	// Create a multi-writer that writes to both the file and stdout
	multiWriter := io.MultiWriter(f, os.Stdout)

	// Set the log output to the multi-writer
	_logrus.SetOutput(multiWriter)

	logsCh = make(chan *LogData, 100)

	// Start a goroutine to handle logs
	go handleLogs()

	return nil
}

func log(ctx context.Context, level, format string, v ...interface{}) {
	// Get function caller info (file and line number)
	_, file, line, _ := runtime.Caller(2)
	//funcName := runtime.FuncForPC(pc).Name()

	// Get values from context
	//serviceName, _ := ctx.Value("service_name").(string)
	logID, _ := ctx.Value("log_id").(string)
	locate, _ := ctx.Value("locate").(string)

	// Create log data
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	data := &LogData{
		Locate: locate,
		//ServiceName:    serviceName,
		LogID:          logID,
		LogLevel:       level,
		LogMessage:     fmt.Sprintf(format, v...),
		ExecutionStack: string(buf),
		FileLine:       fmt.Sprintf("%s:%d", file, line),
		Timestamp:      fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)), // change to millisecond timestamp
	}

	// Convert string level to logrus.Level
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		_logrus.Warn("Failed to parse log level: %v", err)
		return
	}

	// Write to logrus
	_logrus.WithFields(logrus.Fields{
		//"service_name":    serviceName,
		"log_id":    logID,
		"file_line": data.FileLine,
	}).Log(logLevel, data.LogMessage)

	// Write to elasticsearch
	wg.Add(1)
	logsCh <- data
}

func handleLogs() {
	var logs []*LogData
	for {
		select {
		case log := <-logsCh:
			logs = append(logs, log)
			if len(logs) >= batchSize {
				sendToElasticsearch(logs)
				logs = logs[:0]
			}
		case <-time.After(time.Second * 5):
			if len(logs) > 0 {
				sendToElasticsearch(logs)
				logs = logs[:0]
			}
		}
	}
}

func sendToElasticsearch(logs []*LogData) {
	bulkRequest := elasticClient.Bulk()
	for _, log := range logs {
		req := elastic.NewBulkIndexRequest().Index("log_index").Doc(log)
		bulkRequest = bulkRequest.Add(req)
	}

	response, err := bulkRequest.Do(context.Background())
	if err != nil {
		_logrus.Log(logrus.WarnLevel, "Failed to write logs to Elasticsearch: %v", err)
		return
	}

	if response.Errors {
		// 打印错误信息
		for _, item := range response.Items {
			for _, result := range item {
				if result.Error != nil {
					_logrus.Log(logrus.WarnLevel, "Error indexing document: %v", result.Error)
				}
			}
		}
	}

	wg.Done()
}

func Error(ctx context.Context, format string, v ...interface{}) {
	log(ctx, "error", format, v...)
}

func Info(ctx context.Context, format string, v ...interface{}) {
	log(ctx, "info", format, v...)
}

func Warn(ctx context.Context, format string, v ...interface{}) {
	log(ctx, "warn", format, v...)
}

func Close() error {
	close(logsCh)
	wg.Wait()
	return nil
}
