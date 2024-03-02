package services

import (
	"fmt"

	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/valueObject"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	// gunplaRepository repository.GunplaRepository
	// toolRepository   repository.ToolRepository
}

func CreateOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
	// gunplaRepository: gunplaRepository, toolRepository: toolRepository
}

func (s *OrderService) GetAllOrders() ([]*entity.Order, error) {
	return s.orderRepository.GetAllOrders()
}
func (s *OrderService) AddOrder(order restModel.OrderRestModal) (*restModel.OrderRestModal, error) {
	// Update gunpla stock
	var gunpla []valueObject.Product
	var tool []valueObject.Product

	for _, p := range order.Cart {
		if p.Type == "Gunpla" {
			gunpla = append(gunpla, p)
		} else if p.Type == "Tool" {
			tool = append(tool, p)
		}
	}
	fmt.Println(gunpla)
	fmt.Println(tool)
	// if _, err := s.gunplaRepository.UpdateGunplaStock(gunpla); err != nil {
	// 	return nil, fmt.Errorf("failed to update gunpla stock: %v", err)
	// }
	if _, err := s.orderRepository.UpdateOrderStock(order); err != nil {
		return nil, fmt.Errorf("failed to update tool stock: %v", err)
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
