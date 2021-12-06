package chatUsecase

import (
	"chatapp/models"
	"chatapp/chat"
)

type ChatUseCase struct{
	repo chat.Repository
}


func (uc *ChatUseCase) SaveMessage(msg *models.Message) error{
	err := uc.repo.SaveMessage(msg)
	return err
}