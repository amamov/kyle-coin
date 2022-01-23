package restapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string

func Start(portNumber int) {
	port = fmt.Sprintf(":%d", portNumber)

	// handler := http.NewServeMux()
	router := mux.NewRouter()

	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", getDocsController).Methods("GET")
	router.HandleFunc("/blocks", getBlocksController).Methods("GET")
	router.HandleFunc("/blocks", appendBlockController).Methods("POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", getBlockController).Methods("GET")

	fmt.Printf("Listening on http://localhost%s REST API ✨ \n", port)

	log.Fatal(http.ListenAndServe(port, router))
}
