package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ueirt/model"
)


func RecoveryHandler(c *gin.Context, err interface{}) {
	fmt.Print("GOTO RecoveryHandler RecoveryHandler RecoveryHandler")
	appGin := &model.Gin{C: c}
	if err != nil {
		switch err.(type) {
		case model.ApiError:
			apiError := err.(model.ApiError)
			appGin.ToErrorResponse(apiError.Status, apiError.Message, apiError.Err)
		default:
			appGin.ToErrorResponse(http.StatusInternalServerError, "Unexpected error", err)
		}
	}
}
