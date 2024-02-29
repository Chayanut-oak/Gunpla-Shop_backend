package entity

type User struct {
	UserId      string  `json:"userId"`
	Email       string  `json:"email"`
	Password       string  `json:"password"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Role        string  `json:"Role"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}
