package internal

import (
	"flag"
	"fmt"
	"strings"
)
func Head() (err error){
	var numOfLines int
	flag.IntVar(&numOfLines,"-n",10,"specify number of lines to be printed")
	flag.Parse()
    path:= flag.Args()[0]
	content , err := ScanFile(path)
	lines := strings.Split(content, "\n")
	if err != nil {
		return err
	}
	for counter,line:= range lines {
		if counter >= numOfLines {
			break
		}
		fmt.Println(line)
	}
	return nil
}
