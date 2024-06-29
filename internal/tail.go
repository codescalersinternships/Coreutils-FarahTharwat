package internal

import (
	"flag"
	"fmt"
	"strings"
)

func Tail() (err error) {
	var numOfLines int
	flag.IntVar(&numOfLines,"-n",10,"specify number of lines to be printed")
	flag.Parse()
    path:= flag.Args()[0]
	content , err := ScanFile(path)
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
