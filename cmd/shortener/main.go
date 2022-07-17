package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/POST" method="POST">
            <label>Url</label><input type="text" name="Url">
            <input type="submit" value="Сократить">
        </form>
    </body>
</html>`

var ShortUrl = map[string]string{}

func Shortening(Url string) string {
	urlS := make([]byte, 5)
	for i := range urlS {
		urlS[i] = Url[rand.Intn(len(Url))]
	}
	ShortUrl[string(urlS)] = Url
	return string(urlS)
}

func Url(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad URL", 400)
			return
		} else {
			Url := r.FormValue("Url")
			w.WriteHeader(201)
			UrlS := Shortening(Url)
			w.Write([]byte(UrlS))
		}
	default:
		fmt.Fprint(w, form)
	}
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
			UrlG := ShortUrl[q]
			w.Header().Set("location", UrlG)
			w.WriteHeader(307)
		}
	}
}

func main() {
	http.HandleFunc("/POST", Url)
	http.HandleFunc("/GET/", Get)
	http.ListenAndServe(":8080", nil)
}
