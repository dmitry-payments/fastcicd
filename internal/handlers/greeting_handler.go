package handlers

import (
	"encoding/json"
	"fastcicd/internal/database"
	"net/http"
)

type GreetingHandler struct {
	repo *database.GreetingRepo
}

func NewGreetingHandler(repo *database.GreetingRepo) *GreetingHandler {
	return &GreetingHandler{repo: repo}
}

func (h *GreetingHandler) GetGreetings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	greetings, err := h.repo.GetGreetings(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(greetings)
}

func (h *GreetingHandler) AddGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.repo.AddGreeting(ctx, req.Message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Greeting added"))
}
