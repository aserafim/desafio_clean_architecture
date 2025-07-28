# 🧱 Desafio Clean Architecture

Este projeto implementa uma aplicação de gerenciamento de pedidos (`orders`) seguindo os princípios da **Clean Architecture**, com suporte a:

- ✅ REST API
- ✅ gRPC API
- ✅ GraphQL API

---

## 📦 Tecnologias

- Golang
- PostgreSQL
- gRPC
- GraphQL (GQLGen)
- Docker / Docker Compose

---

## 🚀 Como rodar

### 1. Clonar o repositório

```bash
git clone https://github.com/seuusuario/desafio_clean_architecture.git
cd desafio_clean_architecture
````

### 2. Subir a infraestrutura com Docker

```bash
docker-compose up -d
```

Isso irá:

* Subir o banco PostgreSQL (`orders_db`)
* Criar a tabela `orders`
* Inserir pedidos de exemplo (via migrations)

---

## ▶️ Rodar a aplicação

```bash
go run ./cmd/ordersystem
```

---

## 📚 Endpoints

### ✅ REST - Porta `:8080`

| Método | Rota      | Descrição        |
| ------ | --------- | ---------------- |
| GET    | `/orders` | Lista os pedidos |

Exemplo com `curl`:

```bash
curl http://localhost:8080/orders
```

---

### 🔁 GraphQL - Porta `:8080`

* Playground: [http://localhost:8080](http://localhost:8080)

#### 📘 Exemplo de query:

```graphql
query {
  listOrders {
    id
    product
    price
    createdAt
  }
}
```

---

### 📡 gRPC - Porta `:50051`

#### ✅ Listar pedidos

```bash
grpcurl -plaintext \
  -proto internal/infra/grpc/protofiles/order.proto \
  localhost:50051 \
  orderpb.OrderService/ListOrders
```

#### ✅ Criar pedido

```bash
grpcurl -plaintext \
  -d '{"id": "550e8400-e29b-41d4-a716-446655440999", "price": 400, "tax": 40}' \
  -proto internal/infra/grpc/protofiles/order.proto \
  localhost:50051 \
  orderpb.OrderService/CreateOrder
```

---

## 🧪 Testes

Você pode rodar os testes com:

```bash
go test ./...
```

---

## 🗃️ Estrutura do projeto

```text
internal/
  application/      # Casos de uso
  domain/           # Entidades de domínio
  infra/            # Interfaces (http, grpc, graphql)
  usecase/          # Casos de uso específicos

cmd/ordersystem/    # Entrypoint da aplicação
api/                # Arquivos .proto e api.http
graph/              # Schema GraphQL
migrations/         # SQL de criação/inserção
```

---

## ✅ Checklist do desafio

| Item                           | Status      |
| ------------------------------ | ----------- |
| REST: GET `/orders`            | ✅ Concluído |
| gRPC: ListOrders + CreateOrder | ✅ Concluído |
| GraphQL: Query `listOrders`    | ✅ Concluído |
| Docker + banco PostgreSQL      | ✅ Concluído |
| Migrations automáticas         | ✅ Concluído |
| Arquivo `api.http`             | ✅ Concluído |
| README.md com instruções       | ✅ Concluído |

---

## ✍️ Autor

Alefe Abdiel Correia Serafim
