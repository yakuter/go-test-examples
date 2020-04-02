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

	// Status code test
	if w.Code != 404 {
		t.Error("Http test isteği başarısız")
	}

	// Return value test
	if w.Body.String() != "pongd" {
		t.Error("Dönen cevap farklı, test başarısız")
	}

}
