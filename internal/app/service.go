package app

import (
	"context"
	"time"
)

// Service define os métodos principais da aplicação
type Service interface {
	GetDollarQuote(ctx context.Context) (string, error)
	SaveDollarQuote(ctx context.Context, quote string) error
}

// serviceImpl é a implementação concreta do Service
type serviceImpl struct {
	api        APIClient
	repository Repository
}

// APIClient interface para obter a cotação da API externa
type APIClient interface {
	FetchDollarQuote(ctx context.Context) (string, error)
}

// Repository interface para persistir dados no banco
type Repository interface {
	Save(ctx context.Context, quote string) error
}

// NewService cria uma nova instância de Service
func NewService(api APIClient, repository Repository) Service {
	return &serviceImpl{api: api, repository: repository}
}

// GetDollarQuote obtém a cotação do dólar
func (s *serviceImpl) GetDollarQuote(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	return s.api.FetchDollarQuote(ctx)
}

// SaveDollarQuote salva a cotação no repositório
func (s *serviceImpl) SaveDollarQuote(ctx context.Context, quote string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	return s.repository.Save(ctx, quote)
}
