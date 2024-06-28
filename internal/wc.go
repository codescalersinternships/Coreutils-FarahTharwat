package internal

import (
	"fmt"
	"os"
	"strings"
)

func Wc()(lineCount,wordCount,bytesCount int, err error) {
	filepath := os.Args[1]
	content , err := ScanFile(filepath)

	lines := strings.Split(content, "\n")
	lineCount = 0
	wordCount = 0
	bytesCount = 0
	if err != nil {
		fmt.Print(0,0,0)
		return 0, 0, 0 , err
	}
	for _,line:= range lines {
		lineCount++
		bytesCount += len(line)
		wordCount += len(strings.Fields(line))
	}
	if lineCount == 0 {
		fmt.Print(0,0,0)
		return 0, 0, 0,nil
	}
	fmt.Print(lineCount,wordCount,bytesCount+1)
	return lineCount, wordCount, bytesCount+1, nil
}
