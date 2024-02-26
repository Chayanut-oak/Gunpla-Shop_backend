package entity

type Order struct {
	OrderId    string   `json:"orderId"`
	UserId     string   `json:"userId"`
	Items      []string `json:"items"`
	TotalPrice float64  `json:"totalPrice"`
}
