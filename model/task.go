package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"
)

type Task struct {
	gorm.Model
	User            User
	UserID          uint
	Tile            string
	Context         string
	Level           int
	Reminder        string
	Deadline        time.Time
	RepetitionCycle int
	AppendixHash    string `json:"appendix_hash"`
	AppendixName    string `json:"appendix_name"`
	Group           string
	Sort            int
}

const (
	LevelMIN int = 0
	LevelMAX int = 1
	LevelMID int = 2
	LevelAVG int = 3
)

func (t Task) CreateTask(task *Task) string {
	newTask := new(Task)
	info := ""
	err := DB.Create(newTask).Error
	if err != nil {
		info = "create task err: " + err.Error()
		zap.L().Debug(info)
		DB.Rollback()
	}
	zap.L().Debug("create task " + task.Tile)
	return info
}

func (t Task) GetAllTaskByUserID(userID uint) []Task {
	var tasks []Task
	DB.Find(&tasks, Task{UserID: userID})
	return tasks
}

func (t Task) GetAllTaskByUserIDLimit(userID uint, page int, pageSize int) []Task {
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

func (t Task) UpdateTaskByTaskID(task Task) string {
	info := ""
	err := DB.Model(&task).Updates(Task{}).Error
	if err != nil {
		info = "create task err: " + err.Error()
		zap.L().Debug(info)
		DB.Rollback()
	}
	zap.L().Debug("create task " + utils.ToString(task.ID))
	return info
}

func (t Task) ToString() string {
	return fmt.Sprintf("%+v", t)
}
