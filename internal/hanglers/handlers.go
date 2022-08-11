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

func Shortening(Url string) string {
	numbers := "1234567890"
	urlS := make([]byte, 5)
	for i := range urlS {
		urlS[i] = []byte(numbers)[rand.Intn(len(numbers))]
	}
	storage.ShortUrl[string(urlS)] = Url
	return string(urlS)
}

func PostUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	Url := string(body)
	UrlS := "http://localhost:8080/" + Shortening(Url)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(201)
	w.Write([]byte(UrlS))
}

func GetUrl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", 400)
		return
	}
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
}
