package internal

import (
	"fmt"
	"strings"
)

func Cat() (err error) {
	args,err := argsParsing(0)
	if err != nil {
		return err	
	}
	flag,_:= args[0].(bool)
    path,_:= args[1].(string)
	content , err := ScanFile(path)
	lines := strings.Split(content, "\n")
	if err != nil {
		return err
	}
	counter := 0
	for _,line := range lines {
		if flag {
			fmt.Println(counter, line)
			counter++
			continue
		}
		fmt.Println(line)
	}
	return nil
}
