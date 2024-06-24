package main

// import (
// 	"Linux_Commands/internal"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	//Trimming the new line character from the input
// 	input = strings.TrimSpace(input)
// 	// Spliting the input by spaces and appending to an array of strings (slice)
// 	arrayOfStrings := strings.Fields(input)
// 	command_type := arrayOfStrings[0]
// 	switch command_type {
// 	case "head", "tail":
// 		filepath := strings.Trim(arrayOfStrings[2], "\"") // Removing surrounding quotes
// 		num, err := strconv.Atoi(arrayOfStrings[1])
// 		if err != nil {
// 			fmt.Println("Error converting input to integer:", err)
// 			return
// 		}
// 		if command_type == "tail" {
// 			internal.Tail(num, filepath)
// 		} else {
// 			internal.Head(num, filepath)
// 		}
// 	case "cat":
// 		if arrayOfStrings[1] == "-n" {
// 			filepath := strings.Trim(arrayOfStrings[2], "\"") // Removing surrounding quotes
// 			internal.Cat(true, filepath)
// 		} else {
// 			filepath := strings.Trim(arrayOfStrings[1], "\"") // Removing surrounding quotes
// 			internal.Cat(false, filepath)
// 		}
// 	case "echo":
// 		internal.Echo(arrayOfStrings[1:])
// 	case "wc":
// 		filepath := strings.Trim(arrayOfStrings[1], "\"") // Removing surrounding quotes
// 		lines, words, bytes := internal.Wc(filepath)
// 		if lines == 0 {
// 			fmt.Println("Error: Incorrect file path or file does not exist")
// 		} else {
// 			fmt.Println(lines, words, bytes)
// 		}
// 	case "tree":
// 		filepath := strings.Trim(arrayOfStrings[1], "\"") // Removing surrounding quotes
// 		internal.Tree(filepath)
// 	case "true":
// 		internal.True()
// 	case "false":
// 		internal.False()
// 	case "yes":
// 		if len(arrayOfStrings) > 1 {
// 			word := strings.Trim(arrayOfStrings[1], "\"") // Removing surrounding quotes
// 			internal.Yes(word)
// 		} else {
// 			internal.Yes(" ")
// 		}
// 	}

// }
