package database

import (
	"github.com/Masterminds/squirrel"
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

	// check if user already exists
	var existingUserID int
	err := r.db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&existingUserID)
	if err == nil {
		return existingUserID, nil
	}

	q := psql.Insert("users").Columns("username", "address", "pubkey").
		Values(user.Username, user.Address, user.PubKey).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db)

	err = q.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserPubKey(username string) (int, string, error) {
	var id int
	var pubkey string
	q, args, err := psql.Select("id, pubkey").From("users").
		Where(squirrel.Eq{"username": username}).Limit(1).ToSql()

	row := r.db.QueryRow(q, args...)
	if err = row.Scan(&id, &pubkey); err != nil {
		return 0, "", err
	}
	return id, pubkey, err
}

func (r *UserPostgres) GetAllUsers() ([]ecdhsnap.User, error) {
	var users []ecdhsnap.User
	query := "SELECT * FROM users"
	err := r.db.Select(&users, query)
	return users, err
}
