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
	Tasks     []Task
	Dialogues []Dialogue
}

type UserGroup struct {
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupID   uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sort      int
	IsCreator bool
}

func (g Group) Create(group *Group) error {
	err := DB.Create(group).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		DB.Rollback()
		return err
	}
	DB.Commit()
	return nil
}

func (g Group) DeleteByID(groupID uint) error {
	err := DB.Delete(&Group{}, groupID).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		DB.Rollback()
		return err
	}
	DB.Commit()
	return nil
}

func (g Group) Update(group *Group) error {
	err := DB.Model(group).Updates(Group{
		IsShare: group.IsShare,
		Title:   group.Title,
	}).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		DB.Rollback()
		return err
	}
	DB.Commit()
	return nil
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

func (g Group) GetAllGroupByUserIDLimit(userID uint, page int, pageSize int) []Task {
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
	DB.Offset(offset).Find(&shareTask).Limit(pageSize)
	return shareTask
}

func (g Group) ToString() string {
	return fmt.Sprintf("%+v", g)
}
