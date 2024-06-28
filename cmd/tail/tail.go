package main
import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	err := internal.Tail()
	if err != nil {
		log.Fatal(err)
	}
}
