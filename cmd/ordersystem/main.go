package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"desafio_clean_architecture/internal/application/usecase"
	httpHandler "desafio_clean_architecture/internal/infra/http"
	"desafio_clean_architecture/internal/infra/repository"
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

	http.HandleFunc("/orders", handler.ListOrders)

	fmt.Println("REST API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
