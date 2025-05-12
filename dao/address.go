package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(address model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

// GetAddressById 根据id获取Address
func (dao *AddressDao) GetAddressById(id uint) (address *model.Address, err error) {
	address = &model.Address{}
	err = dao.DB.Where("id = ?", id).First(address).Error
	return address, err
}

// GetAddressesByUserId 获取user_id下的所有地址
func (dao *AddressDao) GetAddressesByUserId(uId uint) (address []*model.Address, err error) {
	address = []*model.Address{}
	err = dao.DB.Model(&model.Address{}).Where("user_id = ?", uId).Find(&address).Error
	return address, err
}

// 当address的某些字段为空，是不算做要更新的字段的
func (dao *AddressDao) UpdateAddressByUserId(uId uint, address *model.Address) (err error) {
	err = dao.DB.Where("user_id = ?", uId).Updates(address).Error
	return err
}

func (dao *AddressDao) DeleteAddressByUserId(uId, aId uint) error {
	err := dao.DB.Model(&model.Address{}).
		Where("id= ? AND user_id = ?", aId, uId).
		Delete(&model.Address{}).
		Error
	return err
}
