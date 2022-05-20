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
	GroupID         uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title           string    `gorm:"default:null"`
	Context         string    `gorm:"default:null"`
	Level           int       `gorm:"default:null"`
	Reminder        time.Time `gorm:"default:null"`
	Deadline        time.Time `gorm:"default:null"`
	RepetitionCycle int       `gorm:"default:null"`
	AppendixHash    string    `gorm:"default:null"`
	AppendixName    string    `gorm:"default:null"`
	Sort            int       `gorm:"default:null"`
}

const (
	LevelMIN int = 0
	LevelMAX int = 1
	LevelMID int = 2
	LevelAVG int = 3
)

func (t Task) Create(task *Task) error {

	e := DB.Create(task).Error
	if e != nil {
		zap.L().Debug(e.Error())
		DB.Rollback()
	}
	return e
}

func (t Task) DeleteByID(taskID uint) error {
	e := DB.Delete(&Task{}, taskID).Error
	if e != nil {
		zap.L().Debug(e.Error())
	}
	return e
}

func (t Task) Update(task *Task) error {
	e := DB.Model(&task).Updates(Task{}).Error
	if e != nil {
		zap.L().Debug(e.Error())
		DB.Rollback()
	}
	zap.L().Debug("create task " + utils.ToString(task.ID))
	return e
}

func (t Task) GetAllTaskByGroupID(GroupID uint) ([]Task, error) {
	var tasks []Task
	e := DB.Find(&tasks, Task{GroupID: GroupID}).Error
	return tasks, e
}

func (t Task) GetAllTaskByGroupIDLimit(GroupID uint, page int, pageSize int) []Task {
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
	DB.Offset(offset).Find(&GroupID, Task{GroupID: GroupID}).Limit(pageSize)
	return shareTask
}

func (t Task) ToString() string {
	return fmt.Sprintf("%+v", t)
}
