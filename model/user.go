package model

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Usr   string `gorm:"type:varchar(50);not null;unique;index"`
	Pwd   string `gorm:"type:varchar(256);not null"`
	Phone int64  `gorm:"index"`
}
