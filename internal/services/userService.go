package services

import (
	"go-backend-api/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

func (us *UserService) GetUserNameByID(id int) string {
	// Dummy implementation
	if id == 1 {
		return "John Doe"
	}
	return ""
}
