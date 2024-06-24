package internal

import (
	"fmt"
	"os"
	//"strings"
)

func Echo() (err error) {
	for _, arg := range os.Args[1:] {
		fmt.Print(arg)
		fmt.Print(" ")
	}
	return nil
}