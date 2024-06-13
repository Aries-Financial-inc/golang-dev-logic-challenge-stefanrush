package main

import (
	"log"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server/handlers"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v\n", r.Method, r.URL)
		w.Write([]byte("OK"))
	})
	server.HandleFunc("POST /analysis", handlers.OnPOSTAnalysis)
	server.HandleFunc("GET /analysis/{contract_id}", handlers.OnGETAnalysisResult)
	server.HandleFunc("GET /analysis/{contract_id}/graph", handlers.OnGETAnalysisResultGraph)

	log.Printf("Starting server on http://localhost:8080\n")
	http.ListenAndServe(":8080", server)
}
