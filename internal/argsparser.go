
package internal

import (
	"fmt"
	"os"
	"strconv"
)
func ArgsParsing(prgm int) (args []interface{},err error) {
	// prgm = 0 :  cat , prgm=1 : head or tail
	values := []interface{}{}
	path:=" "
	switch prgm {
		case 0:
			flag:=false
			fmt.Printf("Os.Args: %v\n" , os.Args)
			path=os.Args[1]
			if os.Args[1]=="-n"{
				flag=true
				path=os.Args[2]}
			values = []interface{}{flag, path}
		case 1:
			if len(os.Args)==2{
				linecount:= 10
				path = os.Args[1]
				values = []interface{}{path,linecount}
			}else if os.Args[1]=="-n"{
				linecount ,err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println("Error:", err)
					return nil,err
				}
				path=os.Args[3]
				values = []interface{}{path,linecount}
			}else {
				fmt.Printf("invalid command typed")
				return nil , fmt.Errorf("invalid command typed")
			}
			
	}
	return values,nil	
} 