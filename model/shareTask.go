package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ShareTask struct {
	gorm.Model
	User   User
	UserID uint
	Task   Task
	TaskID uint
}

func (t ShareTask) GetAllShareTaskByUserID(userID uint) []Task {
	var tasks []Task
	DB.Find(&tasks, Task{UserID: userID})
	return tasks
}

func (t ShareTask) GetAllShareTaskByUserIDLimit(userID uint, page int, pageSize int) []Task {
	var shareTask []Task
	if page == 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	DB.Offset(offset).Find(&userID, Task{UserID: userID}).Limit(pageSize)
	return shareTask
}

func (t ShareTask) ToString() string {
	return fmt.Sprintf("%+v", t)
}
