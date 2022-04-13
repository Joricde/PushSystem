package service

import "PushSystem/model"

type MsgService struct {
}

func (m MsgService) GetAllTaskByUserID(userID uint) *[]model.Task {
	t := new(model.Task)
	return t.GetAllTaskByUserID(userID)
}

func (m MsgService) SetTask(task *model.Task) bool {
	return task.UpdateTask(task)
}

func (m MsgService) DeleteTaskByID(taskID uint) bool {
	t := new(model.Task)
	return t.DeleteTaskByTaskID(taskID)
}
