package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeHandler tests the home handler endpoint
func TestHomeHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body contains what we expect
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal response: %v", err)
	}

	// Vérifie seulement que la response contient une structure valide
	if response.Status == "" {
		t.Errorf("expected a non-empty status, got empty")
	}

	if response.Message == "" {
		t.Errorf("expected a non-empty message, got empty")
	}
}

// TestHealthHandler tests the health endpoint
func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response structure
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal response: %v", err)
	}

	// Vérifie seulement que la response contient un status
	if response.Status == "" {
		t.Errorf("expected a non-empty status, got empty")
	}
}

// TestHelloHandler tests the hello endpoint
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello?name=TestUser", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response structure
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal response: %v", err)
	}

	// Vérifie seulement que la response contient un status et un message
	if response.Status == "" {
		t.Errorf("expected a non-empty status, got empty")
	}

	if response.Message == "" {
		t.Errorf("expected a non-empty message, got empty")
	}
}
