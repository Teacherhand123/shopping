package service

import (
	"context"
	"shopping/dao"
	"shopping/pkg/e"
	"shopping/pkg/utils"

	"shopping/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousel, err := carouselDao.ListCarousel()

	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCarousels(carousel), uint(len(carousel)))
}
