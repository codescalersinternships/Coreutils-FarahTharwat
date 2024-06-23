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
	if lineCount == 0 {
		return 0, 0, 0
	}
	return lineCount, wordCount, bytesCount
}
