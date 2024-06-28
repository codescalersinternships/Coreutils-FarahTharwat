package main
import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	err := internal.Yes()
	if err != nil {
		log.Fatal(err)
	}
}
