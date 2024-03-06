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

	// check if pair already exists
	var existingPairID int
	err := r.db.QueryRow("SELECT id FROM pair WHERE (id_user_initiator = $1 AND id_user_partner = $2) OR (id_user_initiator = $2 AND id_user_partner = $1)",
		pair.Id_user_initiator, pair.Id_user_partner).Scan(&existingPairID)
	if err == nil {
		return existingPairID, nil
	}

	var id int
	q := psql.Insert("pair").Columns("id_user_initiator", "id_user_partner").
		Values(pair.Id_user_initiator, pair.Id_user_partner).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(r.db)

	err = q.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PairPostgres) GetRelatedPairs(userID int) ([]ecdhsnap.User, error) {
	var users []ecdhsnap.User
	query, args, err := psql.
		Select("DISTINCT u.id, u.username, u.address, u.pubkey").
		From("users u").
		Join("pair p on u.id = p.id_user_initiator or u.id = p.id_user_partner").
		Where("p.id_user_initiator = ? OR p.id_user_partner = ?", userID, userID).
		ToSql()

	rows, err := r.db.Query(query, args...)
	for rows.Next() {
		var user ecdhsnap.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Address, &user.PubKey); err != nil {
			return users, err
		}
		if user.Id != userID {
			users = append(users, user)
		}

	}
	return users, err
}
