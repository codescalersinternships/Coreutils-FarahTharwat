package commands
import (
	"bufio"
	"fmt"
	"os"
)

func Cat(flag bool , filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Incorrect file path or file does not exist")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		if flag {
			fmt.Println(counter, scanner.Text())
			counter++
			continue
		}
		fmt.Println(scanner.Text())
	}
}
