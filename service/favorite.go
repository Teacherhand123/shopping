package service

import (
	"context"
	"shopping/dao"
	"shopping/model"
	"shopping/pkg/e"
	"shopping/pkg/utils"
	"shopping/serializer"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

// Create
func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.FavoritesExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)

	// 该用户不存在
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	boss, err := userDao.GetUserById(service.BossId)

	// 该老板不存在
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		UserId:    uId,
		User:      *user,
		BossId:    service.BossId,
		Boss:      *boss,
		ProductId: service.ProductId,
		Product:   *product,
	}

	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Show(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success

	favorites, err := favoriteDao.GetFavoritesById(uId)

	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}

// Delete
func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success

	err := favoriteDao.DeleteFavorite(uId, fId)

	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
