package service

import (
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
