package main
import (
	"Linux_Commands/internal"
	"os"
)
func main() {
	_,_,_,err := internal.Wc()
	if err != nil {
		os.Exit(1)}
}
