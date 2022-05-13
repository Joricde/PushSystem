package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Dialogue struct {
	gorm.Model
	GroupID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID  uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Context string
}

func (d Dialogue) CreateDialogue(dialogue *Dialogue) error {
	err := DB.Create(dialogue).Error
	return err
}

func (d Dialogue) DeleteDialogueByID() error {
	e := DB.Delete(d).Error
	return e
}

func (d Dialogue) UpdateDialogue(dialogue *Dialogue) error {
	err := DB.Model(dialogue).Updates(Dialogue{
		Context: dialogue.Context,
	}).Error
	return err
}

func (d Dialogue) GetAllDialogueByGroupID(groupID uint) (*[]Dialogue, error) {
	dialogues := new([]Dialogue)
	e := DB.Model(Group{Model: gorm.Model{ID: groupID}}).
		Association("Dialogues").Find(dialogues)
	return dialogues, e
}

func (d Dialogue) ToString() string {
	return fmt.Sprintf("%+v", d)
}
