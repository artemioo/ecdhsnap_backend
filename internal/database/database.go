package database

import "github.com/jmoiron/sqlx"

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

type Database struct {
	Authorization
	Pair
	Message
}

// конструктор
func NewDatabase(db *sqlx.DB) *Database {
	return &Database{}
}
