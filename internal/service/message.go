package service

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
)

type MessageService struct {
	db database.Message
}

func NewMessageService(db database.Message) *MessageService {
	return &MessageService{db: db}
}

func (s *MessageService) CreateMessage(message ecdhsnap.Message) (int, error) {
	return s.db.CreateMessage(message)
}
