package commands
import (
	"bufio"
	"fmt"
	"os"
)

func Head(num int, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Incorrect file path or file does not exist")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() && counter <= num {
		fmt.Println(scanner.Text())
		counter++
	}
	}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	//Trimming the new line character from the input
// 	input = strings.TrimSpace(input)
// 	// Spliting the input by spaces and appending to an array of strings (slice)
// 	arrayOfStrings := strings.Fields(input)
// 	num, err := strconv.Atoi(arrayOfStrings[0])
// 	if err != nil {
// 		fmt.Println("Error converting input to integer:", err)
// 		return
// 	}
// 	head(num,arrayOfStrings[1])

// }