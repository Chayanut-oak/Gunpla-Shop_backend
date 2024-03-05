package entity

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"
)

type Order struct {
	OrderId        string                `json:"orderId"`
	Email          string                `json:"email"`
	Cart           []valueObject.Product `json:"cart"`
	TotalPrice     int64                 `json:"totalPrice"`
	Status         string                `json:"status"`
	ShippingMethod string                `json:"shippingMethod"`
	OrderDate      string                `json:"orderDate"`
	Address        string                `json:"address"`
}
