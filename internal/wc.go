package internal

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

func Wc()(lineCount,wordCount,bytesCount int, err error) {
	path := os.Args[1]
	content , err := scanFile(path)
	lines := strings.Split(content, "\n")
	wordCount = 0
	bytesCount = 0
	if err != nil {
		fmt.Print(0,0,0)
		return 0, 0, 0 , err
	}
	for _,line:= range lines {
		bytesCount += len(line)+1
		wordCount += len(strings.Fields(line))
	}
	if len(lines)-1 == 0 {
		fmt.Print(0,0,0)
		return 0, 0, 0,nil
	}
	fmt.Println(len(lines)-1,wordCount,bytesCount-1,filepath.Base(path))
	return len(lines)-1, wordCount, bytesCount-1, nil
}
