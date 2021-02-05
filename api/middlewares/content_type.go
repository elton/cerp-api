package middlewares

import "github.com/gin-gonic/gin"

// SetMiddlewareJSON a middleware for JSON responses
func SetMiddlewareJSON() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Next()
	}
}
