package main

import (
	"Linux_Commands/internal"
	"log"
)

func main() {
	err := internal.Cat()
	if err != nil {
		log.Fatal(err)
	}
}

