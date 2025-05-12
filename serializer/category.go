package serializer

import "shopping/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           (*item).ID,
		CategoryName: (*item).CategoryName,
		CreateAt:     (*item).CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.Category) (Categorys []Category) {
	for _, item := range items {
		Category := BuildCategory(item)
		Categorys = append(Categorys, Category)
	}
	return Categorys
}
