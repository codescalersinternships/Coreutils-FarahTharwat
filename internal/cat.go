package internal

import (
	"flag"
	"fmt"
	"strings"
)

func Cat() (err error) {
	var linesNum bool
	flag.BoolVar(&linesNum,"-n",false, "specify if the number of lines to be printed along the text")
	flag.Parse()
    path:= flag.Args()[0]
	content , err := ScanFile(path)
	lines := strings.Split(content, "\n")
	if err != nil {
		return err
	}
	for counter,line := range lines {
		if linesNum {
			fmt.Println(counter, line)
			continue
		}
		fmt.Println(line)
	}
	return nil
}
