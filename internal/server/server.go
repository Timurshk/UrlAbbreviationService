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
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if len(b) == 0 {
			http.Error(w, "the body cannot be an empty", 400)
			return
		}
		type LongURL = string
		Url := string(b)
		UrlS := hanglers.Shortening(Url)
		print(Url, UrlS)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		w.Write([]byte(UrlS))
	} else if r.Method == http.MethodGet {
		q := r.URL.Query().Get("id")
		print(q)
		if q == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		} else {
			UrlG := storage.ShortUrl[q]
			if UrlG == "" {
				http.Error(w, "the body cannot be an empty", 400)
			}
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Location", UrlG)
			w.WriteHeader(307)
		}
	} else {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
}

func Server() {
	http.HandleFunc("/", Url)
	http.ListenAndServe(":8080", nil)
}
