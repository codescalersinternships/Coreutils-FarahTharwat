package commands
import (
	"fmt"
	"os"
	"strings"
)
func Echo(args []string) {
	for _, arg := range args {
	if arg[0] ==  '$' {
		dollarIndex := strings.Index(arg, "$")
		fmt.Println(os.Getenv(arg[dollarIndex+1:]))
		continue
	}
	fmt.Print(string(arg))
	fmt.Print(" ")
		}
	}
