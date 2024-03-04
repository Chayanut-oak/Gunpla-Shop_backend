package entity

type Gunpla struct {
	ProductId   string   `json:"productId"`
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Series      string   `json:"series"`
	Scale       string   `json:"scale"`
	Grade       string   `json:"grade"`
	Price       int64    `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
}
