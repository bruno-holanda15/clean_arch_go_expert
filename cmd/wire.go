//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"clean_arch_desafio/internal/entity"
	"clean_arch_desafio/internal/infra/database"
	"clean_arch_desafio/internal/infra/web"
	"clean_arch_desafio/internal/usecase"
	"clean_arch_desafio/pkg/events"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
