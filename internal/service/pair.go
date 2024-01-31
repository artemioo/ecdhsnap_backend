package service

import (
	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
)

type PairService struct {
	db database.Pair
}

func NewPairService(db database.Pair) *PairService {
	return &PairService{db: db}
}

func (s *PairService) CreatePair(pair ecdhsnap.Pair) (int, error) {
	return s.db.CreatePair(pair)
}
