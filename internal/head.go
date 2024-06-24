package internal

import (
	"fmt"
	"strings"
)
func Head() (err error){
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
	counter := 0
	for _,line:= range lines {
		if counter >= num ||  counter > len(lines) {
			break
		}else{
			fmt.Println(line)
		}
		counter++
	}
	return nil
}
