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
	q := psql.Insert("users").Columns("username", "address", "pubkey").
		Values(user.Name, user.Address, user.PubKey).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db)

	err := q.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserPubKey(id int) (string, error) {
	var pubkey string
	q, args, err := psql.Select("pubkey").From("users").
		Where(squirrel.Eq{"id": id}).Limit(1).ToSql()

	row := r.db.QueryRow(q, args...)
	if err = row.Scan(&pubkey); err != nil {
		return "", err
	}
	return pubkey, err
}
