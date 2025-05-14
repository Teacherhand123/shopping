package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(order).Error
}

// GetOrderById 根据id获取Order
func (dao *OrderDao) GetOrderByIdAnduId(id, uId uint) (order *model.Order, err error) {
	order = &model.Order{}
	err = dao.DB.Where("id = ? AND user_id = ?", id, uId).First(order).Error
	return order, err
}

// GetOrderesByUserId 获取user_id下的所有地址
func (dao *OrderDao) GetOrderesByUserId(uId uint) (order []*model.Order, err error) {
	order = []*model.Order{}
	err = dao.DB.Model(&model.Order{}).Where("user_id = ?", uId).Find(&order).Error
	return order, err
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, err error) {
	err = dao.Model(&model.Order{}).
		Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).
		Find(&order).Error
	return
}

// 当order的某些字段为空，是不算作要更新的字段的
func (dao *OrderDao) UpdateOrderById(oId uint, order *model.Order) (err error) {
	err = dao.DB.Where("id = ?", oId).Updates(order).Error
	return err
}

func (dao *OrderDao) DeleteOrderByUserId(uId, oId uint) error {
	err := dao.DB.Model(&model.Order{}).
		Where("id= ? AND user_id = ?", oId, uId).
		Delete(&model.Order{}).
		Error
	return err
}
