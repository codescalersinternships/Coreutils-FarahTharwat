package internal

import (
	"fmt"
	"strings"
)

func Tail() (err error) {
	args,err := argsParsing(1)
	if err != nil {
		return err	
	}
    path,_:= args[0].(string)
	num:= args[1].(int)
	content , err := ScanFile(path)
	lines := strings.Split(content, "\n")
	if err != nil {
		return err
	}
	LinesCount,_,_,err := Wc()
	if err!=nil{
		return err
	}
	for counter,line:= range lines{
		if ((LinesCount-counter) > num) {
			break 
		}
		fmt.Println(line)
	}
	return nil
	
}
