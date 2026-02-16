package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/database"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/graphql/graph"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/web"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(repository)
	listOrdersUseCase := usecase.NewListOrdersUseCase(repository)

	webOrderHandler := web.NewWebOrderHandler(createOrderUseCase, listOrdersUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /order", webOrderHandler.Create)
	mux.HandleFunc("GET /order", webOrderHandler.List)

	// GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}}))

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	fmt.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
