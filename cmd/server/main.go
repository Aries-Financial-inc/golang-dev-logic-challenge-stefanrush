package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-stefanrush/server"
)

const port = 8080

func main() {
	server := server.NewServer()
	log.Printf("Server listening on http://localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), server)
}
