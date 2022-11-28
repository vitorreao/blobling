package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	rg.POST("/", createHandler(db))
}
