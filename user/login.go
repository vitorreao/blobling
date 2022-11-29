package user

import (
	b64 "encoding/base64"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorreao/blobling/failure"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request loginRequest
		if err := ctx.BindJSON(&request); err != nil {
			failure.BadRequest(ctx, "Could not deserialize json")
			return
		}
		var user User
		res := db.First(&user, "username = ?", request.Username)
		if res.Error != nil {
			log.Printf("Could not retrieve user: %s", res.Error.Error())
			failure.InternalError(ctx, "Could not retrieve user")
			return
		}
		hashBytes, err := b64.StdEncoding.DecodeString(user.PasswordHash)
		if err != nil {
			log.Printf("Could not decode password hash: %s", err.Error())
			failure.InternalError(ctx, "Could not login user")
			return
		}
		compErr := bcrypt.CompareHashAndPassword(hashBytes, []byte(request.Password))
		if compErr != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
