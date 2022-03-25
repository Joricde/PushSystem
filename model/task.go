package model

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	User            User
	UserID          int
	Tile            string
	Context         string
	Level           string
	Reminder        string
	Deadline        time.Time
	RepetitionCycle int
	AppendixHash    string `json:"appendix_hash"`
	AppendixName    string `json:"appendix_name"`
	group           string
}
