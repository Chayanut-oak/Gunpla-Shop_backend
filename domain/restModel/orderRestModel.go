package restModel

import "github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"

type OrderRestModal struct {
	UserId     string                `json:"userId"`
	Cart       []valueObject.Product `json:"cart"`
	TotalPrice float64               `json:"totalPrice"`
	Address    string                `json:"address"`
}
