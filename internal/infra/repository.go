package infra

import (
	"context"
	"database/sql"

	"github.com/robertocorreajr/fullcycle-client-server-api/internal/app"
)

// repository implementa a interface Repository
type repository struct {
	db *sql.DB
}

// NewRepository cria uma nova instância de Repository
func NewRepository(db *sql.DB) app.Repository {
	return &repository{db: db}
}

// Save salva a cotação no banco de dados
func (r *repository) Save(ctx context.Context, quote string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO cotacoes (valor) VALUES (?)", quote)
	return err
}
