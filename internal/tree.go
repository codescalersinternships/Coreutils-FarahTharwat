package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
func Tree () (err error){
	rootpath := "."
	if len(os.Args) > 1 {
		rootpath = os.Args[1]
	}
	err = filepath.Walk(rootpath, FindPath(rootpath))
	if err != nil {
		fmt.Printf("error in the path %v: %v\n", rootpath, err)
		return err
	}
	return nil
}

func FindPath(rootpath string) filepath.WalkFunc {
	// using closure to keep track if it the root . or only the first subdirectory as filepath.Rel gives same depth for both
	isroot := true
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relativePath, err := filepath.Rel(rootpath, path)
		if err != nil {
			return err
		}
		fdepth := len(strings.Split(relativePath, string(os.PathSeparator))) -1
		printing := ""
		if fdepth ==0 && isroot {
			printing = "|-- "
		} else {
			if fdepth == 0{
				printing = strings.Repeat("|   ", 1) + "|-- "
			}else{
				printing = strings.Repeat("|   ", fdepth+1) + "|-- "
			}
		}
		if info.IsDir() {
			if isroot {
				isroot = false
				fmt.Printf("%s%s\n", printing, ".")
			}else{
				fmt.Printf("%s%s\n", printing, info.Name())
			}
		} else {
			fmt.Printf("%s%s\n", printing, info.Name())
		}

		return nil
	}
}