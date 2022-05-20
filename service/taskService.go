package service

import (
	"PushSystem/model"
	"gorm.io/gorm"
	"time"
)

type TaskService struct {
	GroupID         uint
	TaskID          uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Title           string
	Context         string
	Level           int
	Reminder        time.Time
	Deadline        time.Time
	RepetitionCycle int
	AppendixHash    string
	AppendixName    string
	Sort            int
}

var TaskModel = new(model.Task)

func (t TaskService) AddTask(service *TaskService) error {
	task := model.Task{
		GroupID:         service.GroupID,
		Title:           service.Title,
		Context:         service.Context,
		Level:           service.Level,
		Reminder:        service.Reminder,
		Deadline:        service.Reminder,
		RepetitionCycle: service.RepetitionCycle,
		AppendixHash:    service.AppendixHash,
		AppendixName:    service.AppendixName,
		Sort:            service.Sort,
	}
	e := TaskModel.Create(&task)
	if e != nil {
		return e
	}
	return nil
}

func (t TaskService) DeleteTask(taskID uint) error {
	e := TaskModel.DeleteByID(taskID)
	return e
}

func (t TaskService) UpdateTask(service *TaskService) error {
	task := model.Task{
		Model:           gorm.Model{ID: service.TaskID},
		GroupID:         service.GroupID,
		Title:           service.Title,
		Context:         service.Context,
		Level:           service.Level,
		Reminder:        service.Reminder,
		Deadline:        service.Reminder,
		RepetitionCycle: service.RepetitionCycle,
		Sort:            service.Sort,
	}
	e := TaskModel.Update(&task)
	return e
}

func (t TaskService) UpdateTaskAppendix(service *TaskService) error {
	task := model.Task{
		Model:        gorm.Model{ID: service.TaskID},
		GroupID:      service.GroupID,
		AppendixHash: service.AppendixHash,
		AppendixName: service.AppendixName,
	}
	e := TaskModel.Update(&task)
	return e

}

func (t TaskService) GetAllTasksByGroupID(groupID uint) ([]TaskService, error) {
	tasks, e := TaskModel.GetAllTaskByGroupID(groupID)
	var taskService = make([]TaskService, len(tasks))
	if e != nil {
		return nil, e
	}
	for i, task := range tasks {
		taskService[i].GroupID = task.GroupID
		taskService[i].TaskID = task.ID
		taskService[i].CreatedAt = task.CreatedAt
		taskService[i].UpdatedAt = task.UpdatedAt
		taskService[i].Title = task.Title
		taskService[i].Context = task.Context
		taskService[i].Level = task.Level
		taskService[i].Reminder = task.Reminder
		taskService[i].Deadline = task.Deadline
		taskService[i].RepetitionCycle = task.RepetitionCycle
		taskService[i].AppendixHash = task.AppendixHash
		taskService[i].AppendixName = task.AppendixName
		taskService[i].Sort = task.Sort
	}
	return taskService, e
}
