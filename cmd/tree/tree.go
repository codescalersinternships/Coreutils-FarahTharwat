package main
import (
	"Linux_Commands/internal"
	"os"
)
func main() {
	err := internal.Tree()
	if err != nil {
		os.Exit(1)}
}
