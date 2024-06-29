package internal

import (
	"fmt"
	"os"
	"strings"
)

func Cat(showLineNums bool,path string) (err error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	newContent:= strings.ReplaceAll(string(contentBytes),"\r\n","\n")
	lines := strings.Split(newContent, "\n")
	for counter,line := range lines {
		if showLineNums {
			fmt.Println(counter, line)
			continue
		}
		fmt.Println(line)
	}
	return nil
}
