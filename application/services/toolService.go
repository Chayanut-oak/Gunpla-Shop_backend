package services

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type ToolService struct {
	toolRepository repository.ToolRepository
}

func CreateToolService(toolRepository repository.ToolRepository) *ToolService {
	return &ToolService{toolRepository: toolRepository}
}

func (s *ToolService) GetAllTools() ([]*entity.Tool, error) {
	return s.toolRepository.GetAllTools()
}
func (s *ToolService) AddTool(tool restModel.ToolRestModal) (*restModel.ToolRestModal, error) {
	return s.toolRepository.AddTool(tool)
}
func (s *ToolService) UpdateTool(tool entity.Tool) (*entity.Tool, error) {
	return s.toolRepository.UpdateTool(tool)
}

//	func (s *ToolService) UpdateToolStock(tool []valueObject.Product) (string, error) {
//		return s.toolRepository.UpdateToolStock(tool)
//	}
func (s *ToolService) DeleteTool(ProductId string) error {
	return s.toolRepository.DeleteTool(ProductId)
}
