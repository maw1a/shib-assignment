package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "service is up"})
}
