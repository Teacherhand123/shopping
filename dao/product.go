package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}

// GetProductById 根据id获取product
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	product = &model.Product{}
	err = dao.DB.Where("id = ?", id).First(product).Error
	return product, err
}
