/*
Copyright © 2023 Caezarr-OSS

*/
package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Response est une structure standard pour les réponses API
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Health représente l'état de santé de l'API
type Health struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

// NewRouter crée et configure le routeur de l'API
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	
	// Routes API
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/api/v1/hello", helloHandler).Methods("GET")
	
	// Middleware pour logger les requêtes
	r.Use(loggingMiddleware)
	
	return r
}

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request
		// log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Répond avec les endpoints disponibles
func homeHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Status:  "success",
		Message: "Test API Server",
		Data: map[string]interface{}{
			"endpoints": []string{
				"/",
				"/health",
				"/api/v1/hello",
			},
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Répond avec l'état de santé de l'API
func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Status:    "up",
		Version:   "0.1.0",
		Timestamp: "2023-06-15T12:00:00Z",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// Exemple d'endpoint API
func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Status:  "success",
		Message: "Hello from test-api!",
		Data: map[string]string{
			"server": "test-api",
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
