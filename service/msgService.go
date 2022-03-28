package service

import "PushSystem/model"

type MsgService struct {
}

func (m MsgService) GetAllTaskByUserID(userID uint) []model.Task {
	t := new(model.Task)
	return t.GetAllTaskByUserID(userID)
}
