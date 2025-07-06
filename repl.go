package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		cleanedInput := cleanInput(userInput)

		fmt.Println("Your command was:", cleanedInput[0])
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	words := strings.Fields(lowercaseText)
	return words
}
