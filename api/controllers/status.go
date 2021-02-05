package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck health checks
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
