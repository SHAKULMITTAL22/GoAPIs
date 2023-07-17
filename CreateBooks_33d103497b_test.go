package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Book struct {
	Title  string
	Author string
	Year   string
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	// This is a stub for the createBooks handler
	// Replace this with your actual implementation
}

func TestCreateBooks_33d103497b(t *testing.T) {
	book1 := &Book{
		Title:  "Test Book 1",
		Author: "Test Author 1",
		Year:   "2001",
	}

	jsonBook1, _ := json.Marshal(book1)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBook1))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createBooks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	book2 := &Book{
		Title:  "", // Title is missing
		Author: "Test Author 2",
		Year:   "2002",
	}

	jsonBook2, _ := json.Marshal(book2)
	req, err = http.NewRequest("POST", "/books", bytes.NewBuffer(jsonBook2))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
