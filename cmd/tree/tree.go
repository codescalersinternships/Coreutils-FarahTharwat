package main
import (
	"github.com/codescalersinternships/Coreutils-FarahTharwat/internal"
	"log"
)
func main() {
	err := internal.Tree()
	if err != nil {
		log.Fatal(err)
	}
}
