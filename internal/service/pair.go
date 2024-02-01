package service

import (
	"encoding/json"

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

func (s *PairService) GetRelatedPairs(userID int) ([]byte, error) {
	users, err := s.db.GetRelatedPairs(userID)

	userMap := make(map[int]ecdhsnap.User) // Создание мапы с юзерами
	for _, user := range users {
		userMap[user.Id] = user
	}

	jsonArray, err := json.Marshal(userMap) // Преобразование мапы в JSON
	if err != nil {
		return nil, err
	}

	return jsonArray, err
}
