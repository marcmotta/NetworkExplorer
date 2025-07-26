// cmd/networkexplorer/main.go
package main

import (
	"flag"
	"log"
	"os"

	"networkexplorer/internal/networkexplorer"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := networkexplorer.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
