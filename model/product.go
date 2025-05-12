package model

import (
	"shopping/cache"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (Product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(Product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (Product *Product) AddView() {
	// 添加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(Product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Product.ID)))
}
