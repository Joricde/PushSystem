package service

import (
	"PushSystem/model"
	"time"
)

type MessageService struct {
	ID        uint
	UpdatedAt time.Time
	Title     string
	IsShare   bool
	Sort      int
}

var GroupModel = new(model.Group)

func (m MessageService) AddGroup(userID uint, service *MessageService) error {
	group := model.Group{
		Title:   service.Title,
		IsShare: service.IsShare,
	}
	userGroup := model.UserGroup{
		UserID: userID,
		Sort:   service.Sort,
	}
	err := UserModel.AppendGroup(&userGroup, &group)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) DeleteGroupByGroupID(UserID, GroupID uint) error {
	err := UserModel.DeleteGroup(UserID, GroupID)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) SetGroupSort(UserID, GroupID uint, sort int) error {
	userGroup := model.UserGroup{GroupID: GroupID, UserID: UserID, Sort: sort}
	err := UserModel.SetGroupSortByGroupID(&userGroup)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) SetGroupInfo(groupID uint, service *MessageService) error {
	group := new(model.Group)
	group.ID = groupID
	group.Title = service.Title
	group.IsShare = service.IsShare
	err := GroupModel.Update(group)
	if err != nil {
		return err
	}
	return nil
}

//func (m MessageService) GetShareGroupCode(groupID uint) error {
//
//}

func (m MessageService) GetAllGroupsByUserID(userID uint) ([]MessageService, error) {
	g, err := UserModel.GetAllGroupByUserID(userID)
	if err != nil {
		return nil, err
	}
	groupsService := make([]MessageService, len(g))
	for i, group := range g {
		groupsService[i].ID = group.ID
		groupsService[i].Title = group.Title
		groupsService[i].IsShare = group.IsShare
		groupsService[i].UpdatedAt = group.UpdatedAt
	}
	return groupsService, err
}
