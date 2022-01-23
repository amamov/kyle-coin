package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const templateDir string = "explorer/templates/"

var templates *template.Template

func Start(portNumber int) {
	handler := http.NewServeMux()
	port := fmt.Sprintf(":%d", portNumber)

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", HomeController)
	handler.HandleFunc("/block", BlockController)

	fmt.Printf("Listening on http://localhost%s HTML Explorer✨ \n", port)

	log.Fatal(http.ListenAndServe(port, handler))
}
