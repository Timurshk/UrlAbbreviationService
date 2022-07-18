package server

import (
	"github.com/Timurshk/internal/hanglers"
	"github.com/Timurshk/internal/storage"
	"io"
	"net/http"
)

type URL struct {
	URL string `json:"URL"`
}

func Url(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(b) == 0 {
		http.Error(w, "the body cannot be an empty", http.StatusBadRequest)
		return
	}
	type LongURL = string
	Url := string(b)
	UrlS := hanglers.Shortening(Url)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(UrlS))
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	} else {
		q := r.URL.Query().Get("id")
		if q == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		} else {
			UrlG := storage.ShortUrl[q]
			w.Header().Add("Location", UrlG)
			http.Redirect(w, r, UrlG, http.StatusTemporaryRedirect)
		}
	}
}

func Server() {
	http.HandleFunc("/POST", Url)
	http.HandleFunc("/GET/", Get)
	http.ListenAndServe(":8080", nil)
}
