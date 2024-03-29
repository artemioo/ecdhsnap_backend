package service

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
)

type User interface {
	CreateUser(user ecdhsnap.User) (int, error)
	GetUserPubKey(username string) (int, string, error)
	GetAllUsers() ([]byte, error)
}

type Pair interface {
	CreatePair(pair ecdhsnap.Pair) (int, error)
	GetRelatedPairs(userID int) ([]byte, error)
}

type Message interface {
	CreateMessage(message ecdhsnap.Message) (int, error)
	GetRelatedMessages(pairID int) ([]byte, error)
}

type Service struct {
	User
	Pair
	Message
}

// конструктор
func NewService(db *database.Database) *Service {
	return &Service{
		User:    NewUserService(db.User),
		Pair:    NewPairService(db.Pair),
		Message: NewMessageService(db.Message),
	}
}
