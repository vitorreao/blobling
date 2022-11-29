package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&User{})
	rg.POST("/", createHandler(db))
	rg.POST("/login", loginHandler(db))
}
