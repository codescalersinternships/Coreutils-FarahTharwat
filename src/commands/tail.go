package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Tail(num int, path string) {
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
	LinesCount,_,_ := Wc(path)
	counter := 0
	for scanner.Scan() {
		if ((LinesCount-counter) <= num){
			fmt.Println(scanner.Text())
		}
		counter++
	}
	}
