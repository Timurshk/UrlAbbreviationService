package hanglers

import (
	"github.com/Timurshk/internal/storage"
	"math/rand"
)

const endpoint = "http://localhost:8080/"

func Shortening(Url string) string {
	numbers := "1234567890"
	urlS := make([]byte, 5)
	for i := range urlS {
		urlS[i] = []byte(numbers)[rand.Intn(len(numbers))]
	}
	storage.ShortUrl[string(urlS)] = endpoint
	return string(urlS)
}
