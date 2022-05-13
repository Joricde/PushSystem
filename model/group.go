package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Title     string
	IsShare   bool
	IsCreator bool
	Task      []Task
	Dialogue  []Dialogue
}

type UserGroup struct {
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupID   uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sort      int
}

func (g Group) CreateGroup(group *Group) (bool, error) {
	err := DB.Create(group).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		return false, err
	}
	DB.Commit()
	return true, nil
}

func (g Group) DeleteGroupByID(group *Group) (bool, error) {
	err := DB.Delete(group).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		return false, err
	}
	DB.Commit()
	return true, nil
}

func (g Group) UpdateGroup(group *Group) (bool, error) {
	err := DB.Model(group).Updates(Group{
		IsShare:   group.IsShare,
		IsCreator: group.IsCreator,
		Title:     group.Title,
	}).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		return false, err
	}
	DB.Commit()
	return true, nil
}

func (g Group) GetGroupByID(groupID uint) (*Group, error) {
	group := new(Group)
	err := DB.First(group, groupID).Error
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	return group, nil
}

func (g Group) GetAllGroupByUserIDLimit(userID uint, page int, pageSize int) *[]Task {
	shareTask := new([]Task)
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
	DB.Offset(offset).Find(&userID).Limit(pageSize)
	return shareTask
}

func (g Group) ToString() string {
	return fmt.Sprintf("%+v", g)
}
