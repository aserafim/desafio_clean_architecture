package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"

	"desafio_clean_architecture/internal/application/usecase"
	httpHandler "desafio_clean_architecture/internal/infra/http"
	"desafio_clean_architecture/internal/infra/repository"

	// gRPC
	grpcHandler "desafio_clean_architecture/internal/infra/grpc"
	"desafio_clean_architecture/internal/infra/grpc/pb/orderpb"

	// GraphQL
	graph "desafio_clean_architecture/internal/infra/graph"
	"desafio_clean_architecture/internal/infra/graph/generated"

	graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=orders_db sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	repo := repository.NewOrderRepository(db)
	listUC := usecase.NewListOrdersUseCase(repo)
	handler := httpHandler.NewOrderHandler(listUC)

	// REST endpoint
	http.HandleFunc("/orders", handler.ListOrders)

	// GraphQL endpoint
	graphQLSrv := graphqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			OrderRepo: repo,
		},
	}))
	http.Handle("/graphql", graphQLSrv)
	http.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))

	// gRPC server em goroutine
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on gRPC port: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcOrderService := grpcHandler.NewOrderServiceServer(listUC)
		orderpb.RegisterOrderServiceServer(grpcServer, grpcOrderService)

		reflection.Register(grpcServer)

		fmt.Println("gRPC server running at :50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}

	}()

	// HTTP (REST + GraphQL)
	fmt.Println("HTTP server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
