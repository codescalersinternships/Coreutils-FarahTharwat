package internal

import (
	"fmt"
	"strings"
)

func Cat(showLineNums bool,path string) (err error) {
	content , err := scanFile(path)
	lines := strings.Split(content, "\n")
	if err != nil {
		return err
	}
	for counter,line := range lines {
		if showLineNums {
			fmt.Println(counter, line)
			continue
		}
		fmt.Println(line)
	}
	return nil
}
