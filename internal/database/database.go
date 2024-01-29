package database

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/jmoiron/sqlx"
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

type Database struct {
	User
	Pair
	Message
}

// конструктор
func NewDatabase(db *sqlx.DB) *Database {
	return &Database{User: NewUserPostgres(db)}
}
