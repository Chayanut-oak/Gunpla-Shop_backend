package entity
import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/"
)
type Tool struct{
	ID          int      `json:"id"`
    Images      []string `json:"images"`
    Name        string   `json:"name"`
    Type        string   `json:"type"`
    Price       float64  `json:"price"`
    Stock       int      `json:"stock"`
    Description string   `json:"description"`
}