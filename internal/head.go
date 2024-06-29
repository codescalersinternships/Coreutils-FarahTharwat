package internal

import (
	"fmt"
	"os"
	"strings"
)
func Head(numOfLines int , path string) (err error){
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	newContent:= strings.ReplaceAll(string(contentBytes),"\r\n","\n")
	lines := strings.Split(newContent, "\n")
	for counter,line:= range lines {
		if counter >= numOfLines {
			break
		}
		fmt.Println(line)
	}
	return nil
}
