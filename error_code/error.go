package error_code

import "fmt"

type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewStatus(code int, message string) *StatusError {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}

func (e *StatusError) WithErrMsg(msg string) *StatusError {
	return NewStatus(e.Code, e.Message+msg)
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Message)
}

var (
	InternalError = NewStatus(1, "internal error")

	InvalidToken = NewStatus(2, "token is expired, please login again")

	InvalidParam = NewStatus(10000, "invalid param")

	RegisterNameDuplicate = NewStatus(10001, "username is already registered")

	NoPermission = NewStatus(10002, "no permission")

	FileNotEnd = NewStatus(10003, "file not upload ")

	FileExist = NewStatus(10004, "the file chunk exist")
)
