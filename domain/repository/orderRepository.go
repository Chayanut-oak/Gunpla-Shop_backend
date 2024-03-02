package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type OrderRepository interface {
	GetAllOrders() ([]*entity.Order, error)
	AddOrder(restModel.OrderRestModal) (*restModel.OrderRestModal, error)
	UpdateOrder(entity.Order) (*entity.Order, error)
	UpdateOrderStock(restModel.OrderRestModal)(string, error)
	DeleteOrder(string) error
}
