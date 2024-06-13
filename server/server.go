package server

import (
	"log"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/handlers"
)

// NewServer returns a new http.ServeMux with all routes defined
func NewServer() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		w.Write([]byte("OK"))
	})
	server.HandleFunc("POST /analysis", handlers.OnPOSTAnalysis)
	server.HandleFunc("GET /analysis/{contract_id}", handlers.OnGETAnalysisResult)
	server.HandleFunc("GET /analysis/{contract_id}/graph", handlers.OnGETAnalysisResultGraph)

	return server
}
