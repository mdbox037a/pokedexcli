package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := make([]string, 0)
		if input != "" {
			clean = cleanInput(input)
			fmt.Printf("Your command was: %s\n", clean[0])
		}
	}
}
