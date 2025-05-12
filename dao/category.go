package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// GetCategoryById 根据id获取Category
func (dao *CategoryDao) GetCategoryById(id uint) (category *model.Category, err error) {
	category = &model.Category{}
	err = dao.DB.Where("id = ?", id).First(category).Error
	return category, err
}

// ListCategory 获取所有Category
func (dao *CategoryDao) ListCategory() (category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return
}
