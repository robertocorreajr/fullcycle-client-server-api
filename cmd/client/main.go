package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	// Lê URL do servidor da variável de ambiente
	serverURL := os.Getenv("SERVER_URL")

	// Define um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Faz a requisição ao servidor para obter a cotação
	quote, err := obterCotacao(ctx, serverURL)
	if err != nil {
		log.Fatalf("Erro ao obter cotação: %v", err)
	}

	// Salva a cotação no arquivo
	if err := salvarCotacaoNoArquivo("cotacao.txt", quote); err != nil {
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	log.Println("Cotação salva com sucesso:", quote)
}

func obterCotacao(ctx context.Context, serverURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL+"/cotacao", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return "", err
	}

	return cotacao.Bid, nil
}

func salvarCotacaoNoArquivo(filename, cotacao string) error {
	content := fmt.Sprintf("Dólar: %s", cotacao)
	return os.WriteFile(filename, []byte(content), 0644)
}
