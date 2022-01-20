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
	port := fmt.Sprintf(":%d", portNumber)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", HomeController)
	http.HandleFunc("/add", AddBlockController)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
