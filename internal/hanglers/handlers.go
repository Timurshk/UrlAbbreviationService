package hanglers

import (
	"github.com/Timurshk/internal/storage"
	"github.com/julienschmidt/httprouter"
	"io"
	"math/rand"
	"net/http"
)

type URL struct {
	URL string `json:"URL"`
}

func Shortening(URL string) string {
	numbers := "1234567890"
	URLS := make([]byte, 5)
	for i := range URLS {
		URLS[i] = []byte(numbers)[rand.Intn(len(numbers))]
	}
	storage.ShortURL[string(URLS)] = URL
	return string(URLS)
}

func PostURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	URL := string(body)
	URLS := "http://localhost:8080/" + Shortening(URL)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(201)
	w.Write([]byte(URLS))
}

func GetURL(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
	id := params.ByName("id")
	if id == "" {
		http.Error(w, "The query parameter is missing", 400)
		return
	}
	URLG := storage.ShortURL[id]
	if URLG == "" {
		http.Error(w, "the body cannot be an empty", 400)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", URLG)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
