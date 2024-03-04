package restModel

type ToolRestModel struct {
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Price       int64    `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
}
