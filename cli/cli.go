package cli

import (
	"coin_tutorial/explorer"
	"coin_tutorial/rest"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Welcome to COIN")
	fmt.Println("Please use the following flags:\n\n")
	fmt.Println("-port: 		Set the Port of the server\n")
	fmt.Println("-mod: 			Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}
func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose mode")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "explorer":
		explorer.Start(*port)
	default:
		usage()
	}
}
