package main

import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	err := internal.Echo()
	if err != nil {
		log.Fatal(err)
	}
}

