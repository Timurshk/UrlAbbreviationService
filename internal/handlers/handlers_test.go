package handlers

import (
	"github.com/Timurshk/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHandler(t *testing.T) {
	h := New()
	s := storage.MockStorage{}
	s.GenerateMockData()
	h.storage = &s
	type want struct {
		code        int
		response    string
		contentType string
	}
	type request struct {
		method string
		target string
		path   string
	}
	tests := []struct {
		name    string
		want    want
		request request
	}{
		{
			name: "simple test Get handler #1",
			want: want{
				code:        http.StatusTemporaryRedirect,
				response:    "<a href=\"https://go.dev\">Temporary Redirect</a>.\n\n",
				contentType: "text/plain; charset=utf-8",
			},
			request: request{
				method: http.MethodGet,
				target: "https://go.dev/GMWJGSAPGA",
				path:   "/{id}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.request.method, tt.request.target, nil)
			w := httptest.NewRecorder()
			router := chi.NewRouter()
			router.Get(tt.request.path, h.GetURL)
			router.ServeHTTP(w, request)
			response := w.Result()
			defer response.Body.Close()
			assert.Equal(t, tt.want.code, response.StatusCode, "invalid response code")
			_, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	h := New()
	s := storage.MockStorage{}
	s.GenerateMockData()
	h.storage = &s
	type want struct {
		code        int
		response    string
		contentType string
	}
	type request struct {
		method string
		target string
		path   string
	}
	tests := []struct {
		name    string
		want    want
		request request
	}{
		{
			name: "simple test Post handler #1",
			want: want{
				code:        http.StatusCreated,
				response:    "https://go.dev/GMWJGSAPGA",
				contentType: "text/plain; charset=utf-8",
			},
			request: request{
				method: http.MethodPost,
				target: "https://go.dev",
				path:   "/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.request.method, tt.request.target, strings.NewReader(tt.request.target))
			w := httptest.NewRecorder()
			router := chi.NewRouter()
			router.Post(tt.request.path, h.PostURL)
			router.ServeHTTP(w, request)
			response := w.Result()
			defer response.Body.Close()
			assert.Equal(t, tt.want.code, response.StatusCode, "invalid response code")
			assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"), "invalid response Content-Type")
			_, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
