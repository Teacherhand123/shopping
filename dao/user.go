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
