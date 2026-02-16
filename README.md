# Desafio Clean Architecture

Este projeto é um desafio de Clean Architecture utilizando Go, Docker, MySQL, gRPC, GraphQL e SQLC.

## Tecnologias

- Go
- Docker
- MySQL
- RabbitMQ (opcional)
- gRPC
- GraphQL (99designs/gqlgen)
- SQLC

## Como executar

### Com Docker (recomendado)

1. Suba os containers:
```bash
docker compose up -d
```
O MySQL é inicializado com as migrations em `sql/migrations` via `docker-entrypoint-initdb.d`.

2. A aplicação sobe automaticamente no container `app`.

### Local (Go)

1. Suba o banco e o RabbitMQ:
```bash
docker compose up -d mysql rabbitmq
```

2. Execute a aplicação:
```bash
go run cmd/ordersystem/main.go
```

## Portas

- **REST (HTTP)**: 8000
- **GraphQL (HTTP)**: 8080
- **gRPC**: 50051
- **MySQL**: 3306
- **RabbitMQ**: 5672 (AMQP) e 15672 (Management)

## Endpoints

### REST
- `POST /order`: cria uma ordem
- `GET /order`: lista as ordens

### GraphQL
- `mutation createOrder`: cria uma ordem
- `query orders`: lista as ordens

Playground: `http://localhost:8080/`

## Arquivos de Teste

- Utilize o arquivo `api.http` para testar as requisições REST e GraphQL.
