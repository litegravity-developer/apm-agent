package custom_error

import "fmt"

type CustomError struct {
	Code    int    `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (c CustomError) Error() string {
	return c.Message
}

func ValidationError(message string) CustomError {
	if message == "" {
		message = ValidationErrMsg
	}

	error_ := createError(message, 400)
	return error_
}

func MissingFieldError(fieldName string) CustomError {
	message := fmt.Sprintf(MissingField, fieldName)
	return ValidationError(message)
}

func InvalidFieldError(fieldName string) CustomError {
	message := fmt.Sprintf(InvalidField, fieldName)
	return ValidationError(message)
}

func ParseError(message string) CustomError {
	if message == "" {
		message = ValidationErrMsg
	}

	error_ := createError(message, 400)
	return error_
}

func createError(message string, code int) CustomError {
	error_ := CustomError{Code: code, Message: message, Status: ErrStatus}
	return error_
}

func PermissionDenied(message string) CustomError {
	if message == "" {
		message = PermissionDeniedErrMsg
	}
	error_ := createError(message, 403)
	return error_
}
