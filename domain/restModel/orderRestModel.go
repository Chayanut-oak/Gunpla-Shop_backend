package restModel

import "github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"

type OrderRestModel struct {
	Email          string                `json:"email"`
	Cart           []valueObject.Product `json:"cart"`
	TotalPrice     int64                 `json:"totalPrice"`
	ShippingMethod string                `json:"shippingMethod"`
	Address        string                `json:"address"`
	Token          string                `json:"token"`
}
