package service

import (
	"PushSystem/model"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type MessageService struct {
	GroupID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	IsShare   bool
	IsCreator bool
	Sort      int
}

var (
	GroupModel     = new(model.Group)
	UserGroupModel = new(model.UserGroup)
)

func (m MessageService) AddGroup(userID uint, service *MessageService) error {
	group := model.Group{
		Title:   service.Title,
		IsShare: service.IsShare,
	}
	userGroup := model.UserGroup{
		UserID:    userID,
		Sort:      service.Sort,
		IsCreator: true,
	}
	err := GroupModel.AppendGroup(&userGroup, &group)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) DeleteGroupByGroupID(userID, groupID uint) error {
	userGroup := model.UserGroup{
		UserID:  userID,
		GroupID: groupID,
	}
	zap.L().Debug(fmt.Sprint(userGroup))
	err := GroupModel.DeleteGroup(&userGroup)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) SetGroupSort(UserID, GroupID uint, sort int) error {
	userGroup := model.UserGroup{GroupID: GroupID, UserID: UserID, Sort: sort}
	err := UserGroupModel.UpdateGroupSortByGroupID(&userGroup)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) SetGroupShare(groupID uint, isShare bool) (*MessageService, error) {
	group := new(model.Group)
	group.ID = groupID
	group.IsShare = isShare
	e := GroupModel.UpdateShare(group)
	messageService := MessageService{
		GroupID:   group.ID,
		CreatedAt: group.CreatedAt,
		IsShare:   isShare,
	}
	if e != nil {
		return nil, e
	}

	return &messageService, nil
}

func (m MessageService) JoinShareGroup(userID, groupID uint, sort int) (bool, error) {
	isBelongToUser, e := m.IsBelongToUser(userID, groupID)
	if e != nil {
		return false, e
	}
	if !isBelongToUser {
		group := model.UserGroup{UserID: userID, GroupID: groupID, Sort: sort}
		_, err := UserGroupModel.Create(&group)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		return false, nil
	}
}

func (m MessageService) SetGroupTitle(service *MessageService) error {
	group := new(model.Group)
	group.ID = service.GroupID
	group.Title = service.Title
	zap.L().Debug(fmt.Sprint(group))
	err := GroupModel.Update(group)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageService) IsBelongToUser(userID, groupID uint) (bool, error) {
	g := model.UserGroup{}
	userGroup, err := g.RetrieveByUserIDAndGroupID(userID, groupID)
	if err != nil {
		return false, err
	}
	if userGroup.UserID != 0 {
		return true, nil
	}
	return false, nil

}

//func (m MessageService) GetShareGroupCode(groupID uint) error {
//
//}

func (m MessageService) GetAllGroupsByUserID(userID uint) ([]MessageService, error) {
	g, err := GroupModel.GetAllGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}
	groupsService := make([]MessageService, len(g))
	for i, group := range g {
		groupsService[i].GroupID = group.ID
		groupsService[i].UpdatedAt = group.UpdatedAt
		groupsService[i].CreatedAt = group.CreatedAt
		groupsService[i].Title = group.Title
		groupsService[i].IsShare = group.IsShare
		groupsService[i].IsCreator = group.IsCreator
		groupsService[i].Sort = group.Sort
	}
	return groupsService, err
}
