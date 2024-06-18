package commands
import (
	"bufio"
	"fmt"
	"os"
)

func Cat(flag bool , filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Incorrect file path or file does not exist")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		if flag {
			fmt.Println(counter, scanner.Text())
			counter++
			continue
		}
		fmt.Println(scanner.Text())
	}
}

// func main (){
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	//Trimming the new line character from the input
// 	input = strings.TrimSpace(input)
// 	// Spliting the input by spaces and appending to an array of strings (slice)
// 	arrayOfStrings := strings.Fields(input)
// 	if arrayOfStrings[0] == "-n" {
// 		cat(true, arrayOfStrings[1])
// 		return
// 	}
// 	cat(false, arrayOfStrings[0])
// }