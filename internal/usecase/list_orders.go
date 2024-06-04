package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	ListOrders    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	ListOrders events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		ListOrders:    ListOrders,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	var orders []entity.Order
	orders, err := c.OrderRepository.ListOrders();

	if err != nil {
		return nil, err
	}

	var dto []OrderOutputDTO
	for _, order := range orders {
		dto = append(dto, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.ListOrders.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.ListOrders)

	return dto, nil
}
