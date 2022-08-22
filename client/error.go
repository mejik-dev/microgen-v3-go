package client

import (
	"fmt"
	"net/http"
)

type MicrogenError struct {
	Message    string
	Status     int
	StatusText string
}

func (err *MicrogenError) Error() string {
	return fmt.Sprintf("%d %s %s", err.Status, err.StatusText, err.Message)
}

func internalError(message string) *MicrogenError {
	return &MicrogenError{
		Message:    message,
		Status:     http.StatusInternalServerError,
		StatusText: "Internal Server Error",
	}
}
