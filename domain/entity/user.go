package entity

type User struct {
	ToolId      string   `json:"toolId"`
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
}