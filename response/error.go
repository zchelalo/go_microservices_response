package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func errors(msg string, status int) Response {
	return &ErrorResponse{
		Status:  status,
		Message: msg,
	}
}

func InternalServerError(msg string) Response {
	return errors(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errors(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return errors(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return errors(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return errors(msg, http.StatusBadRequest)
}

func (errorResponse *ErrorResponse) StatusCode() int {
	return errorResponse.Status
}

func (errorResponse *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(errorResponse)
}

func (errorResponse *ErrorResponse) Error() string {
	return errorResponse.Message
}

func (errorResponse *ErrorResponse) GetData() interface{} {
	return nil
}
