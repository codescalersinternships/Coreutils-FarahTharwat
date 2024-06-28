package internal

import (
	"fmt"
	"os"
	"strings"
)

func Echo() (err error) {
	fmt.Println(strings.Join(os.Args[1:]," "))
	return nil
}