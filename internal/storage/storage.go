package storage

import (
	"errors"
	"math/rand"
	"sync"
)

type Storage interface {
	Load(sl string) (string, error)
	Store(url string) (sl string)
}

type DB struct {
	data map[string]string
	mu   sync.Mutex
}

func (d *DB) Load(sl string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	link, ok := d.data[sl]
	if !ok {
		return link, errors.New("url not found")
	}
	return link, nil
}

func (d *DB) Store(url string) (sl string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	sl = Shortening()
	d.data[sl] = url
	return
}

func New() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

//type DB struct {
//	mx  sync.Mutex
//	URL map[string]string
//}
//
//func NewDB() *DB {
//	return &DB{
//		URL: make(map[string]string),
//	}
//}
//
//func (d *DB) Load(key string) (string, bool) {
//	d.mx.Lock()
//	defer d.mx.Unlock()
//	val, ok := d.URL[key]
//	return val, ok
//}
//
//func (d *DB) Store(key string, value string) {
//	d.mx.Lock()
//	defer d.mx.Unlock()
//	d.URL[key] = value
//}
//
//type URL struct {
//	URL string `json:"URL"`
//}

func Shortening() string {
	numbers := "1234567890"
	urls := make([]byte, 5)
	for i := range urls {
		urls[i] = []byte(numbers)[rand.Intn(len(numbers))]
	}
	return string(urls)
}
