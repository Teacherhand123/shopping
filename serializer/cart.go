package serializer

import (
	"context"
	"shopping/conf"
	"shopping/dao"
	"shopping/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	Name          string `json:"name"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            (*cart).ID,
		UserId:        (*cart).UserId,
		ProductId:     (*cart).ProductId,
		CreatedAt:     (*cart).CreatedAt.Unix(),
		Num:           int((*cart).Num),
		Name:          (*product).Name,
		MaxNum:        int((*cart).MaxNum),
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + (*product).ImgPath,
		Check:         (*cart).Check,
		DiscountPrice: (*product).DiscountPrice,
		BossId:        (*cart).BossId,
		BossName:      (*boss).UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (Carts []Cart) {
	productMap := map[uint]*model.Product{}
	bossMap := map[uint]*model.User{}
	productDao := dao.NewProductDao(ctx)
	userDao := dao.NewUserDao(ctx)

	for _, item := range items {
		productMap[item.ProductId] = &model.Product{}
		bossMap[item.BossId] = &model.User{}
	}

	// 查找产品，放入map中
	for pId, _ := range productMap {
		product, err := productDao.GetProductById(pId)
		if err != nil {
			continue
		}
		productMap[pId] = product
	}

	// 查找老板，放入map中
	for uId, _ := range bossMap {
		boss, err := userDao.GetUserById(uId)
		if err != nil {
			continue
		}
		bossMap[uId] = boss
	}

	for _, item := range items {
		Cart := BuildCart(item, productMap[item.ProductId], bossMap[item.UserId])
		Carts = append(Carts, Cart)
	}

	return Carts
}
