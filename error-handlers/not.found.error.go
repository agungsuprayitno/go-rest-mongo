package errorhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotfoundError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (handler NotfoundError) SetError(ctx *gin.Context, message string) {
	errorHandler := NotfoundError{
		Status:  http.StatusNotFound,
		Message: message,
		Code:    "not-found",
	}

	ctx.AbortWithStatusJSON(errorHandler.Status, gin.H{"error": errorHandler})
	return
}
