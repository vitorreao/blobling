package failure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailWithBody(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, NewErrorMsg(message))
}

func BadRequest(ctx *gin.Context, message string) {
	FailWithBody(ctx, http.StatusBadRequest, message)
}

func InternalError(ctx *gin.Context, message string) {
	FailWithBody(ctx, http.StatusInternalServerError, message)
}
