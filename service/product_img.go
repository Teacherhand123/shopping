package service

import (
	"context"
	"shopping/dao"
	"shopping/serializer"
	"strconv"
)

type ListProductImg struct{}

func (service *ListProductImg) List(ctx context.Context, pid string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pid)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))

	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
