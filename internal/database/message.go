package database

import (
	"github.com/Masterminds/squirrel"
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/jmoiron/sqlx"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) CreateMessage(message ecdhsnap.Message) (int, error) {
	var id int
	q := psql.Insert("message").Columns("id_pair", "encrypted_message", "sent_at").
		Values(message.Id_pair, message.Encrypted_message, message.Sent_at).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db)

	err := q.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MessagePostgres) GetRelatedMessages(pairID int) ([]ecdhsnap.Message, error) {
	var messages []ecdhsnap.Message
	query, args, err := psql.
		Select("id, id_pair, encrypted_message, sent_at").
		From("message").
		Where(squirrel.Eq{"id_pair": pairID}).
		ToSql()

	rows, err := r.db.Query(query, args...)
	for rows.Next() {
		var message ecdhsnap.Message
		if err := rows.Scan(&message.Id, &message.Id_pair, &message.Encrypted_message, &message.Sent_at); err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, err
}
