package service

import "github.com/artemioo/ecdhsnap_backend/internal/database"

type Authorization interface {
	// CreateUser()
	// GetUser()
	// GeneratePubKey(entropy)
	// GetPubKey()

}

type Pair interface {
	// CreatePair()
	// GetPair()
	// SaveSharedSecret()
}

type Message interface {
	// CreateMessage()
	// GetMessages()
	//
}

type Service struct {
	Authorization
	Pair
	Message
}

// конструктор
func NewService(db *database.Database) *Service {
	return &Service{}
}
