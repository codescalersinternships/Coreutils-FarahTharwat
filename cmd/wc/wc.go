package main
import (
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
	"log"
)
func main() {
	_,_,_,err := internal.Wc()
	if err != nil {
		log.Fatal(err)
	}
}
