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

var books []Book = []Book{
    {ID: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: "1951"},
    {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: "1960"},
    {ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: "1925"},
    {ID: "4", Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez", Year: "1967"},
    {ID: "5", Title: "The Color Purple", Author: "Alice Walker", Year: "1982"},
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "Application/json")

    // Get params
    params := mux.Vars(r)

    // Loop through books and find id
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Book{})

}

func TestGetBook_01d147860b(t *testing.T) {
    // Create a new request
    req, err := http.NewRequest("GET", "/books/1", nil)
    if err != nil {
        t.Error(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Call the handler
    getBook(rr, req)

    // Check the status code
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
    }

    // Check the response body
    var book Book
    err = json.Unmarshal(rr.Body.Bytes(), &book)
    if err != nil {
        t.Error(err)
    }

    // Check the book ID
    if book.ID != "1" {
        t.Errorf("Expected book ID %s, got %s", "1", book.ID)
    }
}

func TestGetBook_NotFound(t *testing.T) {
    // Create a new request
    req, err := http.NewRequest("GET", "/books/999", nil)
    if err != nil {
        t.Error(err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Call the handler
    getBook(rr, req)

    // Check the status code
    if rr.Code != http.StatusNotFound {
        t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
    }

    // Check the response body
    var book Book
    err = json.Unmarshal(rr.Body.Bytes(), &book)
    if err != nil {
        t.Error(err)
    }

    // Check the book ID
    if book.ID != "" {
        t.Errorf("Expected book ID %s, got %s", "", book.ID)
    }
}
