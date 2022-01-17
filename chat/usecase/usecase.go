package chatUsecase

import (
	"chatapp/models"
	"chatapp/chat"
)

type ChatUseCase struct{
	repository chat.Repository
}

func NewChatUseCase(rep chat.Repository) *ChatUseCase{
	return &ChatUseCase{
		repository: rep,
	}
}

func (uc *ChatUseCase) SaveMessage(msg models.Message) error{
	err := uc.repository.SaveMessage(msg)
	return err
}