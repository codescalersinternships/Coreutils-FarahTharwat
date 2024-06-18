package main

import (
	"Linux_Commands/src/commands"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	//Trimming the new line character from the input
	input = strings.TrimSpace(input)
	// Spliting the input by spaces and appending to an array of strings (slice)
	arrayOfStrings := strings.Fields(input)
	command_type := arrayOfStrings[0]
	switch command_type {
	case "head", "tail":
		num, err := strconv.Atoi(arrayOfStrings[1])
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			return
		}
		commands.Head(num, arrayOfStrings[2])
		break
	case "cat":
		if arrayOfStrings[1] == "-n" {
			commands.Cat(true, arrayOfStrings[2])
		} else {
			commands.Cat(false, arrayOfStrings[1])
		}
	case "echo":
		commands.Echo(arrayOfStrings[1:])
		break
	case "wc":
		commands.Wc(arrayOfStrings[1])
		break
	case "tree":
		commands.Tree(arrayOfStrings[1])
	}

}
