package main

import (
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
	"log"
)
func main() {
	err := internal.Echo()
	if err != nil {
		log.Fatal(err)
	}
}

