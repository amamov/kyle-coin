package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/amamov/kyle-coin/explorer"
	"github.com/amamov/kyle-coin/restapi"
)

const (
	REST_API_MODE string = "restapi"
	HTML_MODE     string = "html"
)

func exitProgram() {
	os.Exit(0)
}

func helpDisplayConsole() {
	fmt.Printf("Welcome to Kyle Coin\n\n")
	fmt.Printf("Pleast use the following flags\n\n")
	fmt.Printf("-port=4000:		Set the PORT of the server\n")
	fmt.Printf("-mode=%s:		Start the REST API\n", REST_API_MODE)
	fmt.Printf("-mode=%s:		Start the HTML explorer\n\n", HTML_MODE)
}

func defaultConsole() {
	helpDisplayConsole()
	exitProgram()
}

func Start() {
	// fmt.Println(os.Args)
	// fmt.Println(reflect.TypeOf(os.Args)) // []string

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", REST_API_MODE, fmt.Sprintf("Choose between '%s' and '%s'", REST_API_MODE, HTML_MODE))
	flag.Parse()

	switch *mode {
	case REST_API_MODE:
		restapi.Start(*port)
	case HTML_MODE:
		explorer.Start(*port)
	default:
		defaultConsole()
	}
}
