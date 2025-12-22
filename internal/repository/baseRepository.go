package repository

import "gorm.io/gorm"

// BaseRepository chứa các phương thức CRUD chung
type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll(limit, offset int) ([]T, error) {
	var list []T
	if err := r.db.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
