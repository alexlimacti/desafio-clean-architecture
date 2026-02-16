package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)
`

type CreateOrderParams struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder,
		arg.ID,
		arg.Price,
		arg.Tax,
		arg.FinalPrice,
	)
	return err
}

const listOrders = `-- name: ListOrders :many
SELECT id, price, tax, final_price FROM orders
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Tax,
			&i.FinalPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
