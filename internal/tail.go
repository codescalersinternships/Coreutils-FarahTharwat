package internal

import (
	"fmt"
	"strings"
)

func Tail(numOfLines int , path string) (err error) {
	content , err := scanFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(content, "\n")
	for i:=len(lines)-1 ; i>(len(lines)-numOfLines);i--{ 
		// in case the numoflines entered in console is larger than the actual number of lines in the file
		if (i==0) {
			break 
		}
		fmt.Println(lines[i])
	}

	return nil
}
