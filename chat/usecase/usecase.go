package chatUsecase

import (
	"chatapp/models"
	"chatapp/chat"
)

type ChatUseCase struct{
	repo chat.Repository
}

func NewChatUseCase(rep chat.Repository) *ChatUseCase{
	return &ChatUseCase{
		repo: rep,
	}
}

func (uc *ChatUseCase) SaveMessage(msg *models.Message) error{
	err := uc.repo.SaveMessage(msg)
	return err
}