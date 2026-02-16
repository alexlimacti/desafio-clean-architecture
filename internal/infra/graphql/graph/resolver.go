package graph

import "github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/usecase"

type Resolver struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}
