package model

import (
	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status  int
	Message string
	Err     error
}

func (e ApiError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func (e ApiError) ToErrorResponse() (int, interface{}) {
	return e.Status, gin.H{"message": e.Message, "error": e.Err}
}
