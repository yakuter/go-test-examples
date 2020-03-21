package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body := w.Body.String()

	// Status code test
	if resp.StatusCode != 200 {
		t.Error("Http test isteği başarısız")
	}

	// Return value test
	if body != "pong3" {
		t.Error("Dönen cevap farklı, test başarısız")
	}

}
