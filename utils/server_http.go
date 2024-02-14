package utils

import (
	"net/http"

	"github.com/ArpitChinmay/interview/entities"
	"github.com/gin-gonic/gin"
)

type appHandler func(ctx *gin.Context) *entities.AppResult

func ServeHttp(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx)
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, entities.Response{
				Success: false,
				Message: "Internal Server Error",
				Data:    nil,
			})
		}
		if result.Err == nil {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: true,
				Message: result.Message,
				Data:    result.Data,
			})
		} else {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: false,
				Message: result.Err.Error(),
				Data:    result.Data,
			})
		}
	}
}
