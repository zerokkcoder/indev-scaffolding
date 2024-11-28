package service

import "github.com/zerokkcoder/indevsca/internal/repository"

type UserService interface {
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		service: service,
		repo:    userRepo,
	}
}

type userService struct {
	service *Service
	repo    repository.UserRepository
}
