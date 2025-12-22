package models

import product "go-backend-api/internal/product/domain/model/entity"

var AllModels = []interface{}{
	&User{},
	&Role{},
	&product.ProductModel{},
}
