package main
import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	err := internal.Tree()
	if err != nil {
		log.Fatal(err)
	}
}
