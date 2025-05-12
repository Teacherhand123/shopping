package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// GetFavoriteById 根据user_id获取Favorite
func (dao *FavoriteDao) GetFavoritesById(id uint) (favorite []*model.Favorite, err error) {
	favorite = []*model.Favorite{}
	err = dao.DB.Where("user_id = ?", id).Find(&favorite).Error
	return favorite, err
}

// ListFavorite 获取Favorite
func (dao *FavoriteDao) ListFavorite() (favorite []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Find(&favorite).Error
	return
}

func (dao *FavoriteDao) FavoritesExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("peoduct_id = ? AND user_id = ?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, err
	}

	return true, nil
}

func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(in).Error
}

func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) error {
	return dao.DB.
		Model(&model.Favorite{}).
		Where("user_id = ? AND id = ?", uId, fId).
		Delete(&model.Favorite{}).
		Error
}
