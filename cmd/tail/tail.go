package main
import (
	"Linux_Commands/internal"
	"os"
)
func main() {
	err := internal.Tail()
	if err != nil {
		os.Exit(1)}
}
