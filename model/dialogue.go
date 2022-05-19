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

func (d Dialogue) GetDialogueByID(groupID uint) ([]Dialogue, error) {
	var dialogue []Dialogue
	e := DB.Where("id = ?", groupID).First(&dialogue).Error
	return dialogue, e
}

func (d Dialogue) GetAllDialogueByGroupID(groupID uint) ([]Dialogue, error) {
	var dialogues []Dialogue
	e := DB.Model(Group{Model: gorm.Model{ID: groupID}}).
		Association("Dialogues").Find(&dialogues)
	return dialogues, e
}

func (d Dialogue) ToString() string {
	return fmt.Sprintf("%+v", d)
}
