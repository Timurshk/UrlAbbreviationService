package storage

import (
	"errors"
	"math/rand"
	"sync"
)

type Storage interface {
	Load(sl string) (string, error)
	Store(url string) (sl string, err error)
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

func (d *DB) Store(url string) (sl string, err error) {
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

func Shortening() string {
	numbers := "1234567890"
	urls := make([]byte, 5)
	for i := range urls {
		urls[i] = []byte(numbers)[rand.Intn(len(numbers))]
	}
	return string(urls)
}
