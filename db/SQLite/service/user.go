package service

import (
	"github.com/escalopa/SQLite/model"
)

type userRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(name string) (model.User, error)
}

type UserService struct {
	ur userRepository
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.ur.GetAllUsers()
}

func (s *UserService) CreateUser(name string) (model.User, error) {
	return s.ur.CreateUser(name)
}
