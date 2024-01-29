package database

import (
	"fmt"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user ecdhsnap.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (username, address, pubkey) VALUES ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Address, user.PubKey)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
