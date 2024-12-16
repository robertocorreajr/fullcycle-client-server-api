package infra

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import anônimo para registrar o driver SQLite
)

// NewSQLiteConnection cria uma conexão SQLite
func NewSQLiteConnection(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, valor TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	return db
}
