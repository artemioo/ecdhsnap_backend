package database

import (
	"fmt"
	"reflect"

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
		if err := rows.Scan(&user.Id, &user.Name, &user.Address, &user.PubKey); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	fmt.Println("Тип переменной x:", reflect.TypeOf(users))
	return users, err
}
