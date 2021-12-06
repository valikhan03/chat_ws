package chat

import "chatapp/models"

type Repository interface {
	SaveMessage(msg *models.Message) error
}
