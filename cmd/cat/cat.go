package main

import (
	"Linux_Commands/internal"
	"os"
)

func main() {
	err := internal.Cat()
	if err == nil {
		os.Exit(0)
	}else {
		os.Exit(1)
	}
}

