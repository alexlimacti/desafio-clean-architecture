package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/database"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/graphql/graph"
	grpcinfra "github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/grpc"
	grpcpb "github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/infra/web"
	"github.com/devfullcycle/gointensivo/20/desafio-clean-architecture/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db := mustOpenDB(buildDSN())
	defer db.Close()

	repository := database.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(repository)
	listOrdersUseCase := usecase.NewListOrdersUseCase(repository)

	startGrpcServer(createOrderUseCase, listOrdersUseCase)
	startGraphQLServer(createOrderUseCase, listOrdersUseCase)
	startWebServer(createOrderUseCase, listOrdersUseCase)
}

func buildDSN() string {
	user := envOrDefault("DB_USER", "root")
	password := envOrDefault("DB_PASSWORD", "root")
	host := envOrDefault("DB_HOST", "localhost")
	port := envOrDefault("DB_PORT", "3306")
	name := envOrDefault("DB_NAME", "orders")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustOpenDB(dsn string) *sql.DB {
	var db *sql.DB
	var err error
	for attempt := 1; attempt <= 10; attempt++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			if pingErr := db.Ping(); pingErr == nil {
				return db
			}
		}
		time.Sleep(2 * time.Second)
	}
	panic(err)
}

func startWebServer(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase) {
	webPort := envOrDefault("WEB_SERVER_PORT", "8000")
	webOrderHandler := web.NewWebOrderHandler(createOrderUseCase, listOrdersUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /order", webOrderHandler.Create)
	mux.HandleFunc("GET /order", webOrderHandler.List)

	fmt.Printf("REST server is running on port %s\n", webPort)
	if err := http.ListenAndServe(":"+webPort, mux); err != nil {
		panic(err)
	}
}

func startGraphQLServer(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase) {
	graphqlPort := envOrDefault("GRAPHQL_SERVER_PORT", "8080")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	go func() {
		fmt.Printf("GraphQL server is running on port %s\n", graphqlPort)
		if err := http.ListenAndServe(":"+graphqlPort, mux); err != nil {
			panic(err)
		}
	}()
}

func startGrpcServer(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase) {
	grpcPort := envOrDefault("GRPC_SERVER_PORT", "50051")

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	grpcpb.RegisterOrderServiceServer(server, grpcinfra.NewOrderService(createOrderUseCase, listOrdersUseCase))

	go func() {
		fmt.Printf("gRPC server is running on port %s\n", grpcPort)
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()
}
