package service

import (
	"context"
	"template-service/internal/model"
	"template-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}
