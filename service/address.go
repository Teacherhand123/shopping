package service

import (
	"context"
	"shopping/dao"
	"shopping/model"
	"shopping/pkg/e"

	"shopping/serializer"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

// 创建
func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	address = model.Address{
		UserId:  uId,
		Name:    service.Name,
		Address: service.Address,
		Phone:   service.Phone,
	}

	err := addressDao.CreateAddress(address)
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

// 展示某个aId的地址
func (service *AddressService) Show(ctx context.Context, aId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	address, err := addressDao.GetAddressById(aId)
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
		Data:   serializer.BuildAddress(address),
	}
}

// 获取user_id下的所有地址
func (service *AddressService) Get(ctx context.Context, uId uint) serializer.Response {
	var addresses []*model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	addresses, err := addressDao.GetAddressesByUserId(uId)
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
		Data:   serializer.BuildAddresss(addresses),
	}
}

// 更新uId和aId的地址
func (service *AddressService) Update(ctx context.Context, uId, aId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	address = &model.Address{
		UserId:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}

	err := addressDao.UpdateAddressByUserId(uId, address)
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
		Data:   serializer.BuildAddress(address),
	}
}

// 删除user_id和aid下的地址
func (service *AddressService) Delete(ctx context.Context, uId, aId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)

	err := addressDao.DeleteAddressByUserId(uId, aId)
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
