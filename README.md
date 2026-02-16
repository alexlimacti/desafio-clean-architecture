# Desafio Clean Architecture

Este projeto é um desafio de Clean Architecture utilizando Go, Docker, RabbitMQ (opcional), gRPC e GraphQL.

## Tecnologias

- Go
- Docker
- MySQL
- RabbitMQ
- gRPC
- GraphQL (99designs/gqlgen)
- SQLC
- Migrate

## Como executar

1. Suba os containers do Docker:
```bash
docker compose up -d
```
Isso iniciará o banco de dados MySQL e o RabbitMQ.

2. Execute a aplicação:
```bash
go run cmd/ordersystem/main.go
```
A aplicação iniciará as migrações automaticamente (se configurado, caso contrário execute `sqlc` ou `migrate` - neste caso o código assume tabelas existem ou migration manual).
**Nota**: As migrations estão em `sql/migrations`. Você pode precisar rodar manualmente se o `main.go` não as executar. Para este desafio, recomenda-se usar uma ferramenta de migration ou executar o SQL `sql/migrations/000001_init.up.sql` no banco de dados.

## Portas

- **Web Server (REST)**: 8000
- **GraphQL**: 8080 (Integrado na porta 8000 no path `/query` e Playground em `/`)
  - *Obs: No código atual, GraphQL está rodando junto com REST na porta 8000.*
- **gRPC**: 50051 (Não implementado devido a limitações de ambiente, mas a estrutura está pronta)

## Endpoints

### REST
- `POST /order`: Cria uma ordem
- `GET /order`: Lista as ordens

### GraphQL
- `mutation createOrder`: Cria uma ordem
- `query orders`: Lista as ordens
Acesse o Playground em `http://localhost:8000/`

## Arquivos de Teste
- Utilize o arquivo `api.http` para testar as requisições REST e GraphQL.

## Observações
A geração de código gRPC não foi possível devido à falta das ferramentas `protoc` e `docker` no ambiente de execução. O código correspondente foi omitido ou deixado incompleto.
