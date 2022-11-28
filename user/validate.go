package user

import "github.com/gin-gonic/gin"

func ValidateSession(ctx *gin.Context) {
	ctx.Next()
}
