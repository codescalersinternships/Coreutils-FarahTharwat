package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CheckFilePath(path string) (string,error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return "",err
	}
	return absPath,nil
}

func CreateFile(filename string,path string) (*os.File, error) {
	absPath, err := CheckFilePath(path)
	if err!=nil {
		return nil,err
	}
	file, err := os.Create(absPath+"/"+filename)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Incorrect file path or file does not exist")
		return nil, err
	}
	return file, nil
}

func ScanFile(filepath string) (content string , err error){
	absPath, err := CheckFilePath(filepath)
	if err!=nil {
		return "",err
	}
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Incorrect file path or file does not exist")
		return "", err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(data), nil
}
