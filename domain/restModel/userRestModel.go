package restModel

type UserRestModel struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
