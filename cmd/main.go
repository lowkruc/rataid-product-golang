package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lowkruc/rataid-product-golang.git/internal/handler"
	"github.com/lowkruc/rataid-product-golang.git/internal/repo"
	"github.com/lowkruc/rataid-product-golang.git/internal/usecase"
)

const (
	serverAddress = "127.0.0.1:10001"
)

func main() {
	// TODO: initial config file, inital logging, etc

	// initial database repository
	dbRepo := repo.New()

	// initial usecase
	usecase := usecase.New(dbRepo)

	// initial http handler
	httpHandler := handler.NewHTTPHandler(usecase)

	// start http serve
	err := http.ListenAndServe(serverAddress, newRoutes(httpHandler))
	if err != nil {
		// if err is not null, exit process and send log
		log.Fatalf("error starting http on %s", serverAddress)
	}

	fmt.Printf("server running: %s", serverAddress)
}
