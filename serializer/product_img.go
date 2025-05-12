package serializer

import (
	"shopping/conf"
	"shopping/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: (*item).ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + (*item).ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		productImgs = append(productImgs, BuildProductImg(item))
	}
	return
}
