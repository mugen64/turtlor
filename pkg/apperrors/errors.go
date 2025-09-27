package apperrors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Message     string
	FieldErrors map[string]string
	Errors      []string
	Status      int
	Code        string
}

func NewApiError(message string, fieldErrors map[string]string, errors []string, status int, code string) *ApiError {
	return &ApiError{
		Message:     message,
		FieldErrors: fieldErrors,
		Errors:      errors,
		Status:      status,
		Code:        code,
	}
}

func BadRequest(message string, fieldErrors map[string]string, errors []string, code string) *ApiError {
	return NewApiError(message, fieldErrors, errors, http.StatusBadRequest, code)
}

func NotFound(message string, code string) *ApiError {
	return NewApiError(message, nil, nil, http.StatusNotFound, code)
}

func RouteNotFound(route string) *ApiError {
	return NewApiError(fmt.Sprintf("%s not found", route), nil, nil, http.StatusNotFound, "ROUTE_NOT_FOUND")
}

func MethodNotAllowed(route string, method string) *ApiError {
	return NewApiError(fmt.Sprintf("method %s not allowed for route %s", method, route), nil, nil, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED")
}

func ValidationError(message string, fieldErrors map[string]string, errors []string) *ApiError {
	return NewApiError(message, fieldErrors, errors, http.StatusBadRequest, "VALIDATION_ERROR")
}

func InternalServerError(message string) *ApiError {
	return NewApiError(message, nil, nil, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR")
}

func (e ApiError) Error() string {
	return e.Format()
}

func (e ApiError) String() string {
	return e.Format()
}

func (e ApiError) Format() string {
	return fmt.Sprintf("%s(%d): %s", e.Code, e.Status, e.Message)
}

func (e ApiError) FormatWithFields() string {
	return fmt.Sprintf("%s(%d): %s", e.Code, e.Status, e.Message)
}

func (e ApiError) Is(other error) bool {
	if other == nil {
		return false
	}

	_, ok := other.(ApiError)
	if !ok {
		return false
	}

	return true
}

func GetErrors(err error) []string {
	if err == nil {
		return nil
	}
	switch v := err.(type) {
	case ApiError:
		return v.Errors
	}

	return []string{err.Error()}
}

func GetFieldErrors(err error) map[string]string {
	if err == nil {
		return nil
	}
	switch v := err.(type) {
	case ApiError:
		return v.FieldErrors
	}

	return nil
}
