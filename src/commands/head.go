package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	//"strings"
)
func Head(num int, path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}
	path = absPath
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Incorrect file path or file does not exist")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() && counter < num {
		fmt.Println(scanner.Text())
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
