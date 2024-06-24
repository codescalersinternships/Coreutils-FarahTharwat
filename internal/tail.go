package internal

import (
	"fmt"
	"strings"
)

func Tail() (err error) {
	args,err := ArgsParsing(1)
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
	counter:= 0
	for _,line:= range lines{
		if ((LinesCount-counter) > num) || counter > len(lines){
			break 
		}
		fmt.Println(line)
		counter++
	}
	return nil
}
