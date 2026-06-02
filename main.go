package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Pokedex >")
		input := scanner.Text()
		clean := cleanInput(input)
		fmt.Printf("Your command was: %s", clean[0])
	}
}
