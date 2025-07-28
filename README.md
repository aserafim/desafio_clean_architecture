
# Desafio Clean Architecture - Listagem de Orders

Este projeto implementa a listagem de pedidos utilizando **Clean Architecture** com suporte a REST, gRPC e GraphQL.

---

## 🚀 Tecnologias

- Go
- PostgreSQL (via Docker)
- REST (GET /order)
- gRPC (`ListOrders`)
- GraphQL (`listOrders`)
- gqlgen
- protoc

---

## 📦 Entidade Order

```go
type Order struct {
  ID        string
  Product   string
  Price     float64
  CreatedAt time.Time
}
```

---

## 🔧 Como executar

### 1. Subir o banco de dados
```bash
docker compose up -d
```

Banco de dados:
- host: `localhost`
- port: `5432`
- user: `postgres`
- password: `postgres`
- db: `orders_db`

---

### 2. Executar aplicação Go (REST + gRPC + GraphQL)

```bash
go run ./cmd/ordersystem
```

---

## 🌐 Portas

- REST: `http://localhost:8080/orders`
- gRPC: `localhost:50051`
- GraphQL: (configure endpoint com gqlgen)

---

## 🧪 Testes com HTTP

Veja o arquivo [`api/api.http`](api/api.http) para testar com REST Client.

---

## ⚙️ gRPC

### Gerar arquivos
```bash
protoc --go_out=. --go-grpc_out=. api/proto/order.proto
```

---

## ⚙️ GraphQL

### Gerar arquivos gqlgen
```bash
go run github.com/99designs/gqlgen generate
```
