package serializer

import (
	"context"
	"shopping/conf"
	"shopping/dao"
	"shopping/model"
)

type Favorite struct {
	UserId     uint   `json:"id"`
	ProductId  uint   `json:"product_id"`
	CreateAt   int64  `json:"create_at"`
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
	Title      string `json:"title"`
	Info       string `json:"info"`
	ImgPath    string `json:"img_path"`
	Price      string `json:"price"`
	BossId     uint   `json:"boss_id"`
	Num        int    `json:"num"`
	OnSale     bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserId:     favorite.UserId,
		ProductId:  favorite.ProductId,
		CreateAt:   favorite.CreatedAt.Unix(),
		Name:       product.Name,
		CategoryId: product.CategoryId,
		Title:      product.Title,
		Info:       product.Info,
		ImgPath:    conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:      product.Price,
		BossId:     boss.ID,
		Num:        product.Num,
		OnSale:     product.OnSale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (Favorites []Favorite) {
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
		Favorite := BuildFavorite(item, productMap[item.ProductId], bossMap[item.UserId])
		Favorites = append(Favorites, Favorite)
	}

	return Favorites
}
