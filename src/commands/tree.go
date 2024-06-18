package commands
import (
	"fmt"
	"os"
	"path/filepath"
)
func Tree (rootpath string){
	err:=filepath.Walk(rootpath, FindPath)
	if err != nil {
		fmt.Println("error walking the path %v: %v\n", rootpath, err)
		return
	}
}

func FindPath(path string, info os.FileInfo, err error) error{
	if err != nil {
        return err
    }
    if info.IsDir() {
        fmt.Printf("Directory: %s\n", path)
    } else {
        fmt.Printf("File: %s\n", path)
    }
    return nil
}