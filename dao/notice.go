package dao

import (
	"context"
	"shopping/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	notice = &model.Notice{}
	err = dao.DB.Where("id = ?", id).First(notice).Error
	return notice, err
}
