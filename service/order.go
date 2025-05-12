package service

import (
	"context"
	"fmt"
	"math/rand"
	"shopping/dao"
	"shopping/model"
	"shopping/pkg/e"
	"strconv"
	"time"

	"shopping/serializer"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"BossId" form:"BossId"`
	UserId    uint    `json:"UserId" form:"UserId"`
	OrderNum  int     `json:"OrderNum" form:"OrderNum"`
	Type      int     `json:"Type" form:"Type"`
	model.BasePage
}

// 创建
func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	addressDao := dao.NewAddressDao(ctx)

	order = &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, // 默认未支付
	}

	// 检验地址存不存在
	address, err := addressDao.GetAddressById(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	order.AddressId = address.ID

	// 订单编号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	// 创建订单
	err = orderDao.CreateOrder(order)
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

// 展示userId和oId的订单
func (service *OrderService) Show(ctx context.Context, uId, oId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	order, err := orderDao.GetOrderByIdAnduId(oId, uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 获取地址
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 获取产品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
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
		Data:   serializer.BuildOrder(order, product, address),
	}
}

// 获取user_id下的所有订单
func (service *OrderService) Get(ctx context.Context, uId uint) serializer.Response {
	var orders []*model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	orders, err := orderDao.GetOrderesByUserId(uId)
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
		Data:   serializer.BuildOrders(ctx, orders),
	}
}

// 更新uId和cId的订单
func (service *OrderService) Update(ctx context.Context, uId, cId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	// 只更新数量
	order = &model.Order{
		Num: uint(service.Num),
	}

	err := orderDao.UpdateOrderById(cId, order)
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

// 删除user_id和aid下的订单
func (service *OrderService) Delete(ctx context.Context, uId, cId uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	err := orderDao.DeleteOrderByUserId(uId, cId)
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
