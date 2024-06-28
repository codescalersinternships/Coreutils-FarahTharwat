package main
import (
	"Linux_Commands/internal"
	"log"
)
func main() {
	_,_,_,err := internal.Wc()
	if err != nil {
		log.Fatal(err)
	}
}
