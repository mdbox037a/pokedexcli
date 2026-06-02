package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		clean := make([]string, 0)
		if input != "" {
			clean = cleanInput(input)
			fmt.Printf("Your command was %s\n", clean[0])
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
