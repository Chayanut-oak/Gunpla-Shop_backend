package entity

import "github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"

type Order struct {
	OrderId    string                `json:"orderId"`
	UserId     string                `json:"userId"`
	Cart       []valueObject.Product `json:"cart"`
	TotalPrice float64               `json:"totalPrice"`
	Address    string                `json:"address"`
}
