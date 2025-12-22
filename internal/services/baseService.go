package services

import (
	"go-backend-api/internal/repository"
)

// BaseService CRUD
type BaseService[T any] struct {
	repo *repository.BaseRepository[T]
}

func NewBaseService[T any](repo *repository.BaseRepository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) Create(entity *T) error {
	return s.repo.Create(entity)
}

func (s *BaseService[T]) Update(entity *T) error {
	return s.repo.Update(entity)
}

func (s *BaseService[T]) Delete(entity *T) error {
	return s.repo.Delete(entity)
}

func (s *BaseService[T]) FindByID(id uint) (*T, error) {
	return s.repo.FindByID(id)
}

func (s *BaseService[T]) FindAll(limit, offset int) ([]T, error) {
	return s.repo.FindAll(limit, offset)
}
