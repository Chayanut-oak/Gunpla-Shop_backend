package restModel

type GunplaRestModel struct {
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Series      string   `json:"series"`
	Scale       string   `json:"scale"`
	Grade       string   `json:"grade"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
}
