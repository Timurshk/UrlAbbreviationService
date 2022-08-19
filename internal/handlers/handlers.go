package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Timurshk/internal/storage"
)

const Host = "http://localhost:8080/%v"

type Handler struct {
	storage storage.Storage
}

func New() *Handler {
	return &Handler{
		storage: storage.New(),
	}
}

func (h *Handler) PostURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed!", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(body) == 0 {
		http.Error(w, "the body cannot be an empty", http.StatusBadRequest)
		return
	}
	url := string(body)
	urls, err := h.storage.Store(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	urlf := fmt.Sprintf(Host, urls)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(urlf))
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "The query parameter is missing", http.StatusBadRequest)
		return
	}
	urlg, ok := h.storage.Load(id)
	if ok != nil {
		http.Error(w, "The query parameter is missing", http.StatusBadRequest)
		return
	}
	if urlg == "" {
		http.Error(w, "the body cannot be an empty", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", urlg)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
