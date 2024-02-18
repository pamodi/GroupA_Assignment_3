package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Item handler returned incorrect status code: got %v, expected %v", status, http.StatusOK)
	}
}
