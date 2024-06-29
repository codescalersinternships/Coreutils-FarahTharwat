package internal

import (
	"fmt"
	"os"
	"strings"
)

func Tail(numOfLines int , path string) (err error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	newContent:= strings.ReplaceAll(string(contentBytes),"\r\n","\n")
	lines := strings.Split(newContent, "\n")
	for i:=(len(lines)-numOfLines); i<= len(lines)-1;i++{ 
		// in case the numoflines entered in console is larger than the actual number of lines in the file
		if (i==0) {
			break 
		}
		fmt.Println(lines[i])
	}

	return nil
}
