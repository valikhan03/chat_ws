package chat

import "chatapp/models"

type UseCase interface {
	SaveMessage(msg models.Message) error
}
