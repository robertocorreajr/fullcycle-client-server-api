package main

import (
	"log"
	"net/http"
	"os"

	"github.com/robertocorreajr/fullcycle-client-server-api/internal/app"
	"github.com/robertocorreajr/fullcycle-client-server-api/internal/infra"
	transportHTTP "github.com/robertocorreajr/fullcycle-client-server-api/internal/transport/http"
)

func main() {
	// Lê variáveis de ambiente
	apiURL := os.Getenv("API_URL")
	dbPath := os.Getenv("DB_PATH")
	serverPort := os.Getenv("SERVER_PORT")

	// Inicializa dependências
	db := infra.NewSQLiteConnection(dbPath)
	repository := infra.NewRepository(db)
	apiClient := infra.NewAPIClient(apiURL)
	service := app.NewService(apiClient, repository)
	handler := transportHTTP.NewHandler(service)

	// Configura servidor
	http.HandleFunc("/cotacao", handler.GetDollarQuoteHandler)
	log.Printf("Servidor rodando na porta %s...\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
