package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type OrderService interface {
	GetAllOrders() ([]*entity.Order, error)
	AddOrder(restModel.OrderRestModal) (*restModel.OrderRestModal, error)
	UpdateOrder(entity.Order) (*entity.Order, error)
	DeleteOrder(string) error
}