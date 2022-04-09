package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open example.txt")
		return
	}

	defer file.Close()

	_, err = file.WriteString("Appending some text to example.txt\n")
	file.WriteString("9089898\n")

	if err != nil {
		fmt.Println("Could not write text to example.txt")

	} else {
		fmt.Println("Operation successful! Text has been appended to example.txt")
	}
}
