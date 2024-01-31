package database

import (
	"github.com/Masterminds/squirrel"
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/jmoiron/sqlx"
)

type PairPostgres struct {
	db *sqlx.DB
}

func NewPairPostgres(db *sqlx.DB) *PairPostgres {
	return &PairPostgres{db: db}
}

func (r *PairPostgres) CreatePair(pair ecdhsnap.Pair) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO users (username, address, pubkey) VALUES ($1, $2, $3) RETURNING id")
	q := psql.Insert("pair").Columns("id_user_initiator", "id_user_partner").
		Values(pair.Id_user_initiator, pair.Id_user_partner).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db)

	err := q.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
