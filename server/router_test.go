package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"
)

func TestRouter(t *testing.T) {
	srv := httptest.NewServer(Router())
	defer srv.Close()
	route := func(p string) string {
		u, _ := url.Parse(srv.URL)
		u.Path = path.Join(u.Path, p)
		return u.String()
	}

	tests := []struct {
		name          string
		method        string
		path          string
		body          io.Reader
		retStatusCode int
	}{
		{
			name:          "test /",
			method:        "GET",
			path:          "/",
			retStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			request, err := http.NewRequest(tt.method, route(tt.path), tt.body)
			if err != nil {
				t.Errorf("http.NewRequest(%v, %v, %v) error = %v", tt.method, tt.path, tt.body, err)
			}

			response, err := client.Do(request)
			if err != nil {
				t.Errorf("client.Do(%v) error = %v", request, err)
			}

			if response.StatusCode != tt.retStatusCode {
				t.Errorf("got status code = %v want %v", response.StatusCode, tt.retStatusCode)
			}
		})
	}
}
