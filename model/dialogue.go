package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dialogue struct {
	gorm.Model
	GroupID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID  uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Context string
}

type ServiceDialogue struct {
	ID       uint
	GroupID  uint
	UserID   uint
	Nickname string
	Context  string
}

func (d Dialogue) Create(dialogue *Dialogue) error {
	err := DB.Create(dialogue).Error
	return err
}

func (d Dialogue) DeleteByID(dialogueID uint) error {
	e := DB.Delete(&Dialogue{}, dialogueID).Error
	return e
}

func (d Dialogue) Update(dialogue *Dialogue) error {
	err := DB.Model(dialogue).Updates(Dialogue{
		Context: dialogue.Context,
	}).Error
	return err
}

func (d Dialogue) GetDialogueByID(id uint) ([]Dialogue, error) {
	var dialogue []Dialogue
	e := DB.Where("id = ?", id).First(&dialogue).Error
	return dialogue, e
}

func (d Dialogue) GetAllDialogueByGroupID(groupID uint) ([]ServiceDialogue, error) {
	var s []ServiceDialogue
	//e := DB.Model(Group{Model: gorm.Model{ID: groupID}}).
	//	Association("Dialogues").Find(&s)
	e := DB.Model(&Dialogue{}).
		Select("*").
		Joins("inner join users u on u.id = dialogues.user_id").
		Scan(&s).Error
	zap.L().Debug(fmt.Sprintln(s))
	return s, e
}

func (d Dialogue) ToString() string {
	return fmt.Sprintf("%+v", d)
}
