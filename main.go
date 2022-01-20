package main

import (
	"github.com/amamov/kyle-coin/explorer"
	"github.com/amamov/kyle-coin/restapi"
)

func main() {
	go explorer.Start(3000)
	restapi.Start(4000)
}
