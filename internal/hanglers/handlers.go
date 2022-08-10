package hanglers

import (
	"github.com/Timurshk/internal/storage"
	"math/rand"
)

func Shortening(Url string) string {
	letters := "qwertyuiopasdfghjklzxcvbnm"
	urlS := make([]byte, 5)
	for i := range urlS {
		urlS[i] = []byte(letters)[rand.Intn(len(Url))]
	}
	storage.ShortUrl[string(urlS)] = Url
	return string(urlS)
}
