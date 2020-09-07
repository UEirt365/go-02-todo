package model

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type ApiError struct {
	Status  int
	Message string
	Err     error
}

func (e ApiError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func (g *Gin) ToErrorResponse(httpStatus int, message string, obj interface{}) {
	g.C.JSON(httpStatus, gin.H{"message": message, "error": obj})
}

func (g *Gin) ResponseFromApiError(apiError ApiError) {
	g.C.JSON(apiError.Status, gin.H{"message": apiError.Message, "error": apiError.Err})
}
