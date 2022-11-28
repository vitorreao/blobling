package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:30;not null;uniqueIndex"`
	PasswordHash string `gorm:"size:80;not null"`
}
