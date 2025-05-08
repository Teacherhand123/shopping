package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetCarouselById 根据id获取Carousel
func (dao *CarouselDao) GetCarouselById(id uint) (carousel *model.Carousel, err error) {
	carousel = &model.Carousel{}
	err = dao.DB.Where("id = ?", id).First(carousel).Error
	return carousel, err
}

// ListCarousel 根据id获取Carousel
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	carousel = make([]model.Carousel, 0)
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return carousel, err
}
