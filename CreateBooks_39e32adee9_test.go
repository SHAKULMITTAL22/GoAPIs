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

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   string `json:"year"`
}

var books []Book

func createBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "Application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)

}

func TestCreateBooks_39e32adee9(t *testing.T) {
    // Test case 1: Create a book with valid data
    reqBody := `{"title": "The Lord of the Rings", "author": "J.R.R. Tolkien", "year": "1954"}`
    req, err := http.NewRequest("POST", "/books", strings.NewReader(reqBody))
    if err != nil {
        t.Error(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(createBooks)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d", http.StatusCreated, status)
    }
    var book Book
    err = json.Unmarshal(rr.Body.Bytes(), &book)
    if err != nil {
        t.Error(err)
    }
    if book.ID == "" {
        t.Error("Expected book ID to be set")
    }
    if book.Title != "The Lord of the Rings" {
        t.Errorf("Expected book title to be 'The Lord of the Rings', got '%s'", book.Title)
    }
    if book.Author != "J.R.R. Tolkien" {
        t.Errorf("Expected book author to be 'J.R.R. Tolkien', got '%s'", book.Author)
    }
    if book.Year != "1954" {
        t.Errorf("Expected book year to be '1954', got '%s'", book.Year)
    }

    // Test case 2: Create a book with invalid data
    reqBody = `{"title": "", "author": "", "year": ""}`
    req, err = http.NewRequest("POST", "/books", strings.NewReader(reqBody))
    if err != nil {
        t.Error(err)
    }
    rr = httptest.NewRecorder()
    handler = http.HandlerFunc(createBooks)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, status)
    }
    var apiError APIError
    err = json.Unmarshal(rr.Body.Bytes(), &apiError)
    if err != nil {
        t.Error(err)
    }
    if apiError.Message != "Invalid book data" {
        t.Errorf("Expected error message 'Invalid book data', got '%s'", apiError.Message)
    }
}
