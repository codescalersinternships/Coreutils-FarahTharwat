package internal

import (
	"fmt"
	"strings"
)
func Head(numOfLines int , path string) (err error){
	content , err := scanFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(content, "\n")
	for counter,line:= range lines {
		if counter >= numOfLines {
			break
		}
		fmt.Println(line)
	}
	return nil
}
