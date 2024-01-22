package database

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
func NewDatabase() *Database {
	return &Database{}
}
