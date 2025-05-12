package service

import (
	"context"
	"shopping/dao"
	"shopping/model"
	"shopping/pkg/e"

	"shopping/serializer"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

// 创建
func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	productDao := dao.NewProductDao(ctx)
	userDao := dao.NewUserDao(ctx)

	// 查看产品是否存在
	product, err := productDao.GetProductById(service.ProductId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	cart = &model.Cart{
		UserId:    uId,
		BossId:    service.BossId,
		ProductId: service.ProductId,
	}

	//创建购物车
	err = cartDao.CreateCart(*cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 查看老板是否存在
	boss, err := userDao.GetUserById(service.BossId)
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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

// 展示某个cId的购物车
// func (service *CartService) Show(ctx context.Context, cId uint) serializer.Response {
// 	var cart *model.Cart
// 	code := e.Success
// 	cartDao := dao.NewCartDao(ctx)

// 	cart, err := cartDao.GetCartById(cId)
// 	if err != nil {
// 		code = e.Error
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}

// 	return serializer.Response{
// 		Status: code,
// 		Msg:    e.GetMsg(code),
// 		Data:   serializer.BuildCart(cart),
// 	}
// }

// 获取user_id下的所有购物车
func (service *CartService) Get(ctx context.Context, uId uint) serializer.Response {
	var carts []*model.Cart
	code := e.Success
	cartDao := dao.NewCartDao(ctx)

	carts, err := cartDao.GetCartesByUserId(uId)
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
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

// 更新uId和cId的购物车
func (service *CartService) Update(ctx context.Context, uId, cId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	cartDao := dao.NewCartDao(ctx)

	// 只更新数量
	cart = &model.Cart{
		Num: uint(service.Num),
	}

	err := cartDao.UpdateCartById(cId, cart)
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

// 删除user_id和aid下的购物车
func (service *CartService) Delete(ctx context.Context, uId, cId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)

	err := cartDao.DeleteCartByUserId(uId, cId)
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
