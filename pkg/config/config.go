package config

import (
	"log"
	"os"
)

// Config guarda as configurações globais da aplicação
type Config struct {
	APIURL     string
	DBPath     string
	ServerPort string
}

// LoadConfig carrega as configurações a partir de variáveis de ambiente
func LoadConfig() *Config {
	config := &Config{
		APIURL:     getEnv("API_URL", "https://economia.awesomeapi.com.br/json/last/USD-BRL"),
		DBPath:     getEnv("DB_PATH", "cotacoes.db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	log.Printf("Configuração carregada: %+v\n", config)
	return config
}

// getEnv obtém o valor de uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
