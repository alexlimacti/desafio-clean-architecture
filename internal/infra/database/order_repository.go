package database

import (
	"context"
	"database/sql"

	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/entity"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/database/db"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	queries := db.New(r.Db)
	return queries.CreateOrder(context.Background(), db.CreateOrderParams{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	})
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	queries := db.New(r.Db)
	ordersModel, err := queries.ListOrders(context.Background())
	if err != nil {
		return nil, err
	}
	var orders []entity.Order
	for _, orderModel := range ordersModel {
		orders = append(orders, entity.Order{
			ID:         orderModel.ID,
			Price:      orderModel.Price,
			Tax:        orderModel.Tax,
			FinalPrice: orderModel.FinalPrice,
		})
	}
	return orders, nil
}
