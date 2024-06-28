package main
import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	err := internal.Head()
	if err != nil {
		log.Fatal(err)
	}
}
