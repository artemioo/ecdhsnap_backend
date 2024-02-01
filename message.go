package ecdhsnap

import (
	"time"
)

type Message struct {
	Id                int
	Id_pair           int
	Encrypted_message string
	Sent_at           time.Time
}
