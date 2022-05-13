package service

import "PushSystem/model"

type MessageService struct {
}

func (m MessageService) GetAllTaskByUserID(GroupID uint) *[]model.Task {
	t := new(model.Task)
	return t.GetAllTaskByGroupID(GroupID)
}

func (m MessageService) SetTask(task *model.Task) bool {
	return task.UpdateTask(task)
}

func (m MessageService) DeleteTaskByID(taskID uint) bool {
	t := new(model.Task)
	return t.DeleteTaskByTaskID(taskID)
}

func (m MessageService) DeleteGroupByGroupID() {

}

func (m MessageService) AddGroup() {

}
