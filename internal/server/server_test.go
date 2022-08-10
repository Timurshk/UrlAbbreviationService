package server

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{name: "positive test #1",
			want: want{
				code:        307,
				response:    `{"status":"ok"}`,
				contentType: "application/json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.args.w, tt.args.r)
		})
	}
}

func TestUrl(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Url(tt.args.w, tt.args.r)
		})
	}
}
