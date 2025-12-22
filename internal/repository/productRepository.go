package repository

// import (
// 	"go-backend-api/internal/models"
// 	"gorm.io/gorm"
// )

// // ProductRepository embed BaseRepository[ProductModel]
// type ProductRepository struct {
// 	*BaseRepository[models.ProductModel]
// }

// func NewProductRepository(db *gorm.DB) *ProductRepository {
// 	return &ProductRepository{
// 		BaseRepository: NewBaseRepository[models.ProductModel](db),
// 	}
// }

// // Nếu cần query riêng biệt, thêm method ở đây
// func (r *ProductRepository) FindByName(name string) (*models.ProductModel, error) {
// 	var product models.ProductModel
// 	if err := r.db.Where("name = ?", name).First(&product).Error; err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }
