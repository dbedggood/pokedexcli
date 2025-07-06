package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    func() error { return commandHelp(commands) },
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) == 0 {
			continue
		}
		firstWord := words[0]

		command, exists := commands[firstWord]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Printf("Error executing command\n")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	words := strings.Fields(lowercaseText)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}
