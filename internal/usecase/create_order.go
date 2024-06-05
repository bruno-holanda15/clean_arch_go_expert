package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/event"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	event := event.NewOrderCreated()
	event.SetPayload(dto)
	c.EventDispatcher.Dispatch(event)

	return dto, nil
}
