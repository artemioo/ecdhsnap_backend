package service

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
)

type User interface {
	CreateUser(user ecdhsnap.User) (int, error)
	// GetUser()
	// GetPubKey()

}

type Pair interface {
	// CreatePair()
	// GetPair()
}

type Message interface {
	// CreateMessage()
	// GetMessages()
}

type Service struct {
	User
	Pair
	Message
}

// конструктор
func NewService(db *database.Database) *Service {
	return &Service{User: NewUserService(db.User)}
}
