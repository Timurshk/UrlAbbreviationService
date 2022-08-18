package storage

import (
	"errors"
)

type MockStorage DB

var data = map[string]string{
	"12345": "https://yandex.ru",
}

func (s *MockStorage) GenerateMockData() {
	for _, v := range data {
		s.Store(v)
	}
}

func (s *MockStorage) Load(urls string) (string, error) {
	link, ok := s.data[urls]
	if !ok {
		return link, errors.New("url not found")
	}

	return link, nil
}

func (s *MockStorage) Store(url string) (urls string) {
	urls = Shortening()

	if s.data == nil {
		s.data = make(map[string]string)
	}

	s.data[urls] = url

	return
}
