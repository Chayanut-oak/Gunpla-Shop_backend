package restModel

type GunplaRestModel struct {
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Series      string   `json:"series"`
	Brand       string   `json:"brand"`
	Scale       string   `json:"scale"`
	Grade       string   `json:"grade"`
	Price       int64    `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
}
