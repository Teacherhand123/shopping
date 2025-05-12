package serializer

import (
	"context"
	"shopping/conf"
	"shopping/dao"
	"shopping/model"
)

type Order struct {
	Id            uint    `json:"id"`
	OrderNum      uint64  `json:"order_num"`
	CreateAt      int64   `json:"create_at"`
	UpdatedAt     int64   `json:"updated_at"`
	UserId        uint    `json:"user_id"`
	ProductId     uint    `json:"product_id"`
	BossId        uint    `json:"boss_id"`
	Num           int     `json:"num"`
	AddressName   string  `json:"address_name"`
	AddressPhone  string  `json:"address_phone"`
	Address       string  `json:"address"`
	Type          uint    `json:"type"`
	ProductName   string  `json:"product_name"`
	ImgPath       string  `json:"img_path"`
	DiscountPrice float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		Id:            order.ID,
		OrderNum:      order.OrderNum,
		CreateAt:      order.CreatedAt.Unix(),
		UpdatedAt:     order.CreatedAt.Unix(),
		UserId:        order.UserId,
		ProductId:     order.ProductId,
		BossId:        order.BossId,
		Num:           order.Num,
		AddressName:   address.Name,
		AddressPhone:  address.Phone,
		Address:       address.Address,
		Type:          order.Type,
		ProductName:   product.Name,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice: order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (Orders []Order) {
	productMap := map[uint]*model.Product{}
	addressMap := map[uint]*model.Address{}
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)

	for _, item := range items {
		productMap[item.ProductId] = &model.Product{}
		addressMap[item.AddressId] = &model.Address{}
	}

	// 查找产品，放入map中
	for pId, _ := range productMap {
		product, err := productDao.GetProductById(pId)
		if err != nil {
			continue
		}
		productMap[pId] = product
	}

	// 查找地址，放入map中
	for aId, _ := range addressMap {
		address, err := addressDao.GetAddressById(aId)
		if err != nil {
			continue
		}
		addressMap[aId] = address
	}

	for _, item := range items {
		Order := BuildOrder(item, productMap[item.ProductId], addressMap[item.UserId])
		Orders = append(Orders, Order)
	}

	return Orders
}
