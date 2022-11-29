package user

type User struct {
	ID           uint   `gorm:"primaryKey;not null;autoIncrement"`
	CreatedAt    string `gorm:"size:20;not null"`
	UpdatedAt    string `gorm:"size:20"`
	Username     string `gorm:"size:30;not null;uniqueIndex"`
	PasswordHash string `gorm:"size:80;not null"`
}
