package main
import (
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
	"log"
)
func main() {
	err := internal.Yes()
	if err != nil {
		log.Fatal(err)
	}
}
