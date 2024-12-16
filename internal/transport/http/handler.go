package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/robertocorreajr/fullcycle-client-server-api/internal/app"
)

// Handler gerencia as requisições HTTP
type Handler struct {
	service app.Service
}

// NewHandler cria uma nova instância de Handler
func NewHandler(service app.Service) *Handler {
	return &Handler{service: service}
}

// GetDollarQuoteHandler lida com a requisição para obter a cotação
func (h *Handler) GetDollarQuoteHandler(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetDollarQuote(r.Context())
	if err != nil {
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		log.Println("Erro ao obter cotação:", err)
		return
	}

	if err := h.service.SaveDollarQuote(r.Context(), quote); err != nil {
		http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
		log.Println("Erro ao salvar cotação:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": quote})
}
