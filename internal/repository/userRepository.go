package repository

import "fmt"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindUserByID(id int) (string, error) {
	// Dummy implementation
	if id == 1 {
		return "John Doe", nil
	}
	return "", fmt.Errorf("User not found")
}
