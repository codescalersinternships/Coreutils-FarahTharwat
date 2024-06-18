package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Wc(filepath string) (int, int, int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Incorrect file path or file does not exist")
		return 0, 0, 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	wordCount := 0
	bytesCount := 0
	for scanner.Scan() {
		lineCount++
		bytesCount += len(scanner.Text())
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}
	fmt.Println(lineCount, wordCount, bytesCount)
	return lineCount, wordCount, bytesCount
}

// func main () {
// 	reader:= bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	//Trimming the new line character from the input
// 	input = strings.TrimSpace(input)
// 	wc(input)
// }
