package server

import (
	"github.com/Timurshk/internal/hanglers"
	"github.com/Timurshk/internal/storage"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type URL struct {
	URL string `json:"URL"`
}

func PostUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if len(body) == 0 {
			http.Error(w, "the body cannot be an empty", 400)
			return
		}
		Url := string(body)
		UrlS := "http://localhost:8080/" + hanglers.Shortening(Url)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		w.Write([]byte(UrlS))
	} else {
		http.Error(w, "Only POST requests are allowed!", 400)
		return
	}
}

func GetUrl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method == http.MethodGet {
		id := params.ByName("id")
		if id == "" {
			http.Error(w, "The query parameter is missing", 400)
			return
		}
		UrlG := storage.ShortUrl[id]
		if UrlG == "" {
			http.Error(w, "the body cannot be an empty", 400)
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", UrlG)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
}

func Server() {
	router := httprouter.New()
	router.POST("/", PostUrl)
	router.GET("/:id", GetUrl)
	http.ListenAndServe("localhost:8080", router)
}
