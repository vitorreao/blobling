package user

import (
	b64 "encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vitorreao/blobling/failure"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type newUserRequest struct {
	Username          string `json:"username"`
	PlainTextPassword string `json:"password"`
}

type newUserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

func createHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request newUserRequest
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, failure.Msg("Could not deserialize json"))
			return
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.PlainTextPassword), 14)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, failure.Msg("Could not hash the password"))
			return
		}
		user := newUser(request.Username, bytes)
		db.Create(user)
		ctx.JSON(http.StatusCreated, newResponse(user))
	}
}

func newUser(usr string, pw []byte) *User {
	pwHash := b64.StdEncoding.EncodeToString(pw)
	return &User{Username: usr, PasswordHash: pwHash}
}

func newResponse(user *User) *newUserResponse {
	return &newUserResponse{user.Username, user.CreatedAt}
}
