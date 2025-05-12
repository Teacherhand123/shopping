package service

import (
	"context"
	"shopping/dao"
	"shopping/pkg/e"
	"shopping/pkg/utils"
	"shopping/serializer"
)

type CategoryService struct{}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()

	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))
}
