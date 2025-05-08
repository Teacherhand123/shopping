package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据User判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	user = &model.User{}
	err = dao.DB.Where("user_name = ?", userName).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false, nil // 用户不存在
	}

	return user, true, err // 用户存在
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

// GetUserById 根据id获取user
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	user = &model.User{}
	err = dao.DB.Where("id = ?", id).First(user).Error
	return user, err
}

func (dao *UserDao) UpdateUserById(id uint, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", id).Updates(user).Error
	return
}
