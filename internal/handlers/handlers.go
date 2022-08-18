package handlers

import (
	"fmt"
	"github.com/Timurshk/internal/storage"
	"github.com/go-chi/chi/v5"
)
import (
	"io"
	"net/http"
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
		http.Error(w, "Only POST requests are allowed!", 400)
		return
	}
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if len(body) == 0 {
		http.Error(w, "the body cannot be an empty", 400)
		return
	}
	url := string(body)
	urls := fmt.Sprintf(Host, h.storage.Store(url))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(201)
	w.Write([]byte(urls))
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "The query parameter is missing", 400)
		return
	}
	urlg, ok := h.storage.Load(id)
	if ok != nil {
		http.Error(w, "The query parameter is missing", 400)
		return
	}
	if urlg == "" {
		http.Error(w, "the body cannot be an empty", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", urlg)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
