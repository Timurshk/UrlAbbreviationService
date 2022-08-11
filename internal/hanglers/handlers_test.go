package hanglers

import (
	"bytes"
	"github.com/Timurshk/internal/storage"
	"github.com/go-playground/assert/v2"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUrl(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name    string
		request string
		long    string
		id      string
		want    want
	}{
		{
			name:    "Positive test",
			request: "http://localhost:8080/123",
			long:    "https://yandex.ru",
			id:      "123",
			want: want{
				contentType: "text/plain",
				statusCode:  307,
			},
		},
		{
			name:    "Negative test",
			request: "http://localhost:8080/456",
			long:    "",
			id:      "456",
			want: want{
				contentType: "text/plain",
				statusCode:  400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				storage.ShortUrl[tt.id] = tt.long
				router := httprouter.New()
				router.GET("/:id", GetUrl)
				req := httptest.NewRequest(http.MethodGet, tt.request, nil)
				rr := httptest.NewRecorder()
				router.ServeHTTP(rr, req)
				result := rr.Result()
				defer result.Body.Close()
				assert.Equal(t, tt.want.statusCode, result.StatusCode)
			}
		})
	}
}

func TestPostUrl(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		request string
		body    string
		want    want
	}{
		{
			name:    "Positive test",
			request: "/",
			body:    "https://yandex.ru",
			want: want{
				statusCode: 201,
			},
		},
		{
			name:    "Negative test",
			request: "/",
			body:    "",
			want: want{
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.POST(tt.request, PostUrl)
			req := httptest.NewRequest(http.MethodPost, tt.request, bytes.NewBufferString(tt.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			result := rr.Result()
			defer result.Body.Close()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}

func TestShortening(t *testing.T) {
	print()
}
