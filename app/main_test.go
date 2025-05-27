package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Hello World!!"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("unexpected body: got %v, want to contain %v", rr.Body.String(), expected)
	}
}