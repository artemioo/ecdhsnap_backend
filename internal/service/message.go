package service

import (
	"encoding/json"

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

func (s *MessageService) GetRelatedMessages(PairID int) ([]byte, error) {

	messages, err := s.db.GetRelatedMessages(PairID)

	MessagesMap := make(map[int]ecdhsnap.Message) // Создание мапы с сообщениями
	for _, message := range messages {
		MessagesMap[message.Id] = message
	}

	jsonArray, err := json.Marshal(MessagesMap) // Преобразую мапу в JSON
	if err != nil {
		return nil, err
	}

	return jsonArray, err
}
