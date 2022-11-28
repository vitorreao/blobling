package user

import (
	b64 "encoding/base64"
	"log"
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
			failure.BadRequest(ctx, "Could not deserialize json")
			return
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.PlainTextPassword), 14)
		if err != nil {
			failure.InternalError(ctx, "Could not hash the password")
			return
		}
		user := newUser(request.Username, bytes)
		res := db.Create(user)
		if res.Error != nil {
			log.Printf("Error saving user to db: %s", res.Error.Error())
			failure.InternalError(ctx, "Error saving user to DB")
			return
		}
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
