package service

import (
	"context"
	"rest-go-api/internal/models"
	"rest-go-api/internal/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Get(cxt context.Context, id int) ([]models.User, error) {
	if id<1 {
		return nil, nil
	}
	return s.repository.Get(cxt, id)
}
