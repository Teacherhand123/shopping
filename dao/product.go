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

// condition map[string]interface{} 这会生成 SQL 查询条件 WHERE category_id = ”
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// 如果 PageSize = 10，表示每次查询最多返回 10 条记录。
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

// 当使用 LIKE 查询时，如果通配符 % 出现在字符串的开头（如 "%keyword%"），数据库无法利用索引进行优化查询。 "%"+info+"%"
// 这是因为 % 表示任意字符，数据库需要逐行扫描表中的数据以匹配条件。
// 分页查询本身不会导致全表扫描，但如果 WHERE 条件已经触发了全表扫描，分页只是在扫描结果中截取一部分数据。
// 如果查询频繁，可以为 title 和 info 字段创建全文索引（Full-Text Index），以提高查询效率。
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {

	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? or info LIKE ?", "%"+info+"%").
		Count(&count).Error

	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? or info LIKE ?", "%"+info+"%").
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).
		Find(&products).
		Error

	// 如果查询频繁，可以为 title 和 info 字段创建全文索引（Full-Text Index），以提高查询效率。
	// 在 MySQL 中，可以使用以下语句创建全文索引：
	// ALTER TABLE products ADD FULLTEXT(title, info);
	// 这是优化后的代码，还没使用
	// err = dao.DB.Model(&model.Product{}).
	// Where("MATCH(title, info) AGAINST(?)", info).
	// Offset((page.PageNum - 1) * (page.PageSize)).
	// Limit(page.PageSize).
	// Find(&products).Error
	return
}

// func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
// 	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).Error
// 	return
// }
