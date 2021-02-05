package responses

import (
	"github.com/gin-gonic/gin"
)

// ResultJSON show the result of responses using JSON.
func ResultJSON(ctx *gin.Context, statusCode int, data interface{}, err error) {
	if err != nil {
		ctx.JSON(statusCode, gin.H{"status": statusCode, "data": nil, "error": err.Error()})
		return
	}
	ctx.JSON(statusCode, gin.H{"status": statusCode, "data": data, "error": nil})
}
