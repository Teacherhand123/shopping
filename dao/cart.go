package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(cart model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&cart).Error
}

// GetCartById 根据id获取Cart
func (dao *CartDao) GetCartById(id uint) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	err = dao.DB.Where("id = ?", id).First(cart).Error
	return cart, err
}

// GetCartesByUserId 获取user_id下的所有地址
func (dao *CartDao) GetCartesByUserId(uId uint) (cart []*model.Cart, err error) {
	cart = []*model.Cart{}
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&cart).Error
	return cart, err
}

// 当cart的某些字段为空，是不算作要更新的字段的
func (dao *CartDao) UpdateCartById(cId uint, cart *model.Cart) (err error) {
	err = dao.DB.Where("id = ?", cId).Updates(cart).Error
	return err
}

func (dao *CartDao) DeleteCartByUserId(uId, cId uint) error {
	err := dao.DB.Model(&model.Cart{}).
		Where("id= ? AND user_id = ?", cId, uId).
		Delete(&model.Cart{}).
		Error
	return err
}
