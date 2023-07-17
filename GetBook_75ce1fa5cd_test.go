package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetBook_75ce1fa5cd is a test function for the getBook function
func TestGetBook_75ce1fa5cd(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBook)

	// We create a mock router and set our handler function to be called when the "/book/{id}" route is hit
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", handler)

	// We then serve our handler with our test request and recorder
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `{"ID":"1","Isbn":"isbn1","Title":"book1","Author":{"Firstname":"author1","Lastname":"lastname1"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Test case where book is not found
	req, err = http.NewRequest("GET", "/book/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected = `{"ID":"","Isbn":"","Title":"","Author":{"Firstname":"","Lastname":""}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
