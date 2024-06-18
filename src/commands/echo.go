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

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	//Trimming the new line character from the input
// 	input = strings.TrimSpace(input)
// 	// Spliting the input by spaces and appending to an array of strings (slice)
// 	arrayOfStrings := strings.Fields(input)
// 	echo(arrayOfStrings)
// }