package usecase

import (
	"clean_arch_desafio/internal/entity"
	"clean_arch_desafio/internal/event"
	"clean_arch_desafio/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	var orders []entity.Order
	orders, err := c.OrderRepository.ListOrders()

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
	event := event.NewOrdersListed()
	event.SetPayload(dto)
	c.EventDispatcher.Dispatch(event)

	return dto, nil
}
