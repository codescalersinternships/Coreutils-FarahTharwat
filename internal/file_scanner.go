package internal

import (
	"fmt"
	"io"
	"os"
)

func scanFile(filepath string) (content string , err error){
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr,"Incorrect file path or file does not exist")
		return "", err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(data), nil
}
