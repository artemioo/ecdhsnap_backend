package service

import (
	"encoding/json"

	ecdhsnap "github.com/artemioo/ecdhsnap_backend"
	"github.com/artemioo/ecdhsnap_backend/internal/database"
)

type UserService struct {
	db database.User
}

func NewUserService(db database.User) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user ecdhsnap.User) (int, error) {
	return s.db.CreateUser(user)
}

func (s *UserService) GetUserPubKey(username string) (int, string, error) {
	return s.db.GetUserPubKey(username)
}

func (s *UserService) GetAllUsers() ([]byte, error) {
	users, err := s.db.GetAllUsers()

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
