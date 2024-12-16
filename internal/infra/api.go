package infra

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/robertocorreajr/fullcycle-client-server-api/internal/app"
)

// apiClient implementa a interface APIClient
type apiClient struct {
	baseURL string
	client  *http.Client
}

// NewAPIClient cria uma nova instância de APIClient
func NewAPIClient(baseURL string) app.APIClient {
	return &apiClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 200 * time.Millisecond},
	}
}

// FetchDollarQuote busca a cotação do dólar na API externa
func (a *apiClient) FetchDollarQuote(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.baseURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response struct {
		USDBRL app.DollarQuoteDTO `json:"USDBRL"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.USDBRL.Bid == "" {
		return "", errors.New("valor bid vazio")
	}

	return response.USDBRL.Bid, nil
}
