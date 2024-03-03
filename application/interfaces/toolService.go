package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type ToolService interface {
	GetAllTools() ([]*entity.Tool, error)
	AddTool(restModel.ToolRestModel) (*restModel.ToolRestModel, error)
	UpdateTool(entity.Tool) (*entity.Tool, error)
	// UpdateToolStock([]valueObject.Product) (string, error)
	DeleteTool(string) error
}
