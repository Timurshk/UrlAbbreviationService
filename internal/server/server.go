package server

import (
	"github.com/Timurshk/internal/hanglers"
	storage "github.com/Timurshk/internal/storage"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type URL struct {
	URL string `json:"URL"`
}

func Url(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		asd := "localhost:8080/" + UrlS
		w.Write([]byte(asd))
	} else if r.Method == http.MethodGet {
		q := params.ByName("id")
		//q := r.URL.Query().Get("id")
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
			w.WriteHeader(http.StatusTemporaryRedirect)
		}
	} else {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
}

func Server() {
	router := httprouter.New()
	router.POST("/", Url)
	router.GET("/:id", Url)
	http.ListenAndServe(":8080", router)
}
