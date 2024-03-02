package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type ToolRepository interface {
	GetAllTools() ([]*entity.Tool, error)
	AddTool(restModel.ToolRestModal) (*restModel.ToolRestModal, error)
	UpdateTool(entity.Tool) (*entity.Tool, error)
	// UpdateToolStock([]valueObject.Product) (string, error)
	DeleteTool(string) error
}
