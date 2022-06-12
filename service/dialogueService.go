package service

import (
	"PushSystem/model"
	"gorm.io/gorm"
)

type DialogueService struct {
	ID       uint
	GroupID  uint
	UserID   uint
	Nickname string
	Context  string
}

var DialogueModel = new(model.Dialogue)

func (d DialogueService) AddDialogue(service DialogueService) error {
	var dialogue = model.Dialogue{
		//Model: gorm.Model{ID: service.ID},
		GroupID: service.GroupID,
		UserID:  service.UserID,
		Context: service.Context,
	}
	err := dialogue.Create(&dialogue)
	return err
}

func (d DialogueService) UpdateContext(service DialogueService) error {
	var dialogue = model.Dialogue{
		Model:   gorm.Model{ID: service.ID},
		Context: service.Context,
	}
	e := dialogue.Update(&dialogue)
	return e
}

func (d DialogueService) GetDialogueByGroupID(groupID uint) ([]DialogueService, error) {
	dialogues, e := DialogueModel.GetDialogueByID(groupID)
	if e != nil {
		return nil, e
	}
	var services []DialogueService
	for i, dialogue := range dialogues {
		services[i].ID = dialogue.ID
		services[i].UserID = dialogue.UserID
		services[i].GroupID = dialogue.GroupID
		services[i].Context = dialogue.Context
	}
	return services, nil
}

func (d DialogueService) DeleteDialogueByID(dialogueID uint) error {
	e := DialogueModel.DeleteByID(dialogueID)
	return e
}
