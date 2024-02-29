package services

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type OrderService struct {
	orderRepository  repository.OrderRepository
	gunplaRepository repository.GunplaRepository
}

func CreateOrderService(orderRepository repository.OrderRepository, gunplaRepository repository.GunplaRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository, gunplaRepository: gunplaRepository}
}

func (s *OrderService) GetAllOrders() ([]*entity.Order, error) {
	return s.orderRepository.GetAllOrders()
}
func (s *OrderService) AddOrder(order restModel.OrderRestModal) (*restModel.OrderRestModal, error) {
	// Update gunpla stock
	if _, err := s.gunplaRepository.UpdateGunplaStock(order); err != nil {
		return nil, fmt.Errorf("failed to update gunpla stock: %v", err)
	}

	// Add order
	addedOrder, err := s.orderRepository.AddOrder(order)
	if err != nil {
		// Rollback gunpla stock update if order addition fails
		// You need to implement a rollback mechanism in the repository method
		// or directly in the service if needed
		// rollbackErr := s.gunplaRepository.RollbackGunplaStockUpdate(order)
		return nil, fmt.Errorf("failed to add order: %v", err)
	}

	return addedOrder, nil
}
func (s *OrderService) UpdateOrder(order entity.Order) (*entity.Order, error) {
	return s.orderRepository.UpdateOrder(order)
}
func (s *OrderService) DeleteOrder(OrderId string) error {
	return s.orderRepository.DeleteOrder(OrderId)
}
