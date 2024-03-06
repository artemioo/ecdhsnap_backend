package database

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user ecdhsnap.User) (int, error)
	GetUserPubKey(username string) (int, string, error)
	GetAllUsers() ([]ecdhsnap.User, error)
}

type Pair interface {
	CreatePair(pair ecdhsnap.Pair) (int, error)
	GetRelatedPairs(userID int) ([]ecdhsnap.User, error)
}

type Message interface {
	CreateMessage(message ecdhsnap.Message) (int, error)
	GetRelatedMessages(pairID int) ([]ecdhsnap.Message, error)
}

type Database struct {
	User
	Pair
	Message
}

// конструктор
func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		User:    NewUserPostgres(db),
		Pair:    NewPairPostgres(db),
		Message: NewMessagePostgres(db),
	}
}
