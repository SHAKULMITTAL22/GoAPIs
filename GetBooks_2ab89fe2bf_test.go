package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "testing"
)

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "Application/json")
    json.NewEncoder(w).Encode(books)
}

func TestGetBooks(t *testing.T) {
    // Test case 1: Get all books
    req, err := http.NewRequest("GET", "/books", nil)
    if err != nil {
        t.Error(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getBooks)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("status code should be 200, got %d", status)
    }

    // Test case 2: Get a specific book
    bookId := rand.Intn(len(books))
    req, err = http.NewRequest("GET", "/books/"+strconv.Itoa(bookId), nil)
    if err != nil {
        t.Error(err)
    }

    rr = httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("status code should be 200, got %d", status)
    }

    // Test case 3: Get a non-existent book
    bookId = len(books) + 1
    req, err = http.NewRequest("GET", "/books/"+strconv.Itoa(bookId), nil)
    if err != nil {
        t.Error(err)
    }

    rr = httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("status code should be 404, got %d", status)
    }
}
