package model

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	gorm.Model
	Title     string
	IsShare   bool `gorm:"default false"`
	Tasks     []Task
	Dialogues []Dialogue
}

type UserGroup struct {
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupID   uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sort      int
	IsCreator bool `gorm:"default true"`
}

type ServiceGroup struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	IsShare   bool
	IsCreator bool
	Sort      int
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
	err := DB.Model(group).Update("title", group.Title).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		DB.Rollback()
		return err
	}
	return nil
}

func (g Group) UpdateShare(group *Group) error {
	err := DB.Model(group).Update("is_share", group.IsShare).Error
	if err != nil {
		zap.L().Error("create group error :" + err.Error())
		DB.Rollback()
		return err
	}
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

func (g Group) AppendGroup(userGroup *UserGroup, group *Group) error {
	e := DB.Create(&group).Error
	if e != nil {
		return e
	}
	userGroup.GroupID = group.ID
	e = DB.Create(userGroup).Error
	if e != nil {
		return e
	}
	zap.L().Debug(fmt.Sprintln(userGroup))
	return e
}

func (g Group) DeleteGroup(userGroup *UserGroup) error {
	e := DB.Where("user_id = ? and group_id = ?",
		userGroup.UserID, userGroup.GroupID).
		Delete(userGroup).Error
	if e != nil {
		return e
	}
	return e
}

func (g UserGroup) UpdateGroupSortByGroupID(userGroup *UserGroup) error {
	e := DB.Model(userGroup).
		Where("group_id = ? and user_id = ?",
			userGroup.GroupID, userGroup.UserID).
		Update("sort = ?", userGroup.Sort).Error
	return e
}

func (g Group) GetAllGroupsByUserID(userID uint) ([]ServiceGroup, error) {
	var serviceGroups []ServiceGroup
	e := DB.Model(&Group{}).
		Select("*").
		Joins("inner join user_groups ug on ug.group_id = groups.id").
		Where("ug.user_id = ? and ug.deleted_at is null", userID).
		Scan(&serviceGroups).Error
	zap.L().Debug(fmt.Sprintln(serviceGroups))
	return serviceGroups, e
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

func (g UserGroup) RetrieveByUserID(userID uint) ([]UserGroup, error) {
	var userGroup []UserGroup
	e := DB.Find(&userGroup).Error
	if e != nil {
		return nil, e
	}
	return userGroup, nil

}

func (g UserGroup) Create(group *UserGroup) (*UserGroup, error) {
	e := DB.Create(&group).Error
	if e != nil {
		return nil, e
	}
	return group, nil
}

func (g UserGroup) RetrieveByUserIDAndGroupID(userID, groupID uint) (*UserGroup, error) {
	userGroup := new(UserGroup)
	e := DB.Where("user_id = ? and group_id = ?",
		userID, groupID).
		First(userGroup).Error
	zap.L().Debug(fmt.Sprintln(userGroup))
	if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return userGroup, e
	}
	return userGroup, nil
}

func (g Group) ToString() string {
	return fmt.Sprintf("%+v", g)
}
