package serializer

import "shopping/model"

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:       item.ID,
		UserId:   item.UserId,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		Seen:     true,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresss(items []*model.Address) (addresss []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresss = append(addresss, address)
	}
	return addresss
}
