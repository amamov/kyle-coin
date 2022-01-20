package restapi

import (
	"fmt"
	"log"
	"net/http"
)

var port string

func Start(portNumber int) {
	port = fmt.Sprintf(":%d", portNumber)
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
