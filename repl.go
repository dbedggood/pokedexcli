package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/dbedggood/pokedexcli/internal/pokeapi"
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
	commands["map"] = cliCommand{
		name:        "map",
		description: "Display names of next 20 areas",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display names of previous 20 areas",
		callback:    commandMapBack,
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

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

var nextUrl string
var prevUrl string

func commandMap() error {
	if err := fetchAndDisplayAreas(nextUrl); err != nil {
		return err
	}
	return nil
}

func commandMapBack() error {
	if err := fetchAndDisplayAreas(prevUrl); err != nil {
		return err
	}
	return nil
}

func fetchAndDisplayAreas(url string) error {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	locationAreaResponse := LocationAreaResponse{}
	if err := pokeapi.Fetch(url, &locationAreaResponse); err != nil {
		return err
	}

	for _, area := range locationAreaResponse.Results {
		fmt.Println(area.Name)
	}

	nextUrl = locationAreaResponse.Next
	prevUrl = locationAreaResponse.Previous
	return nil
}
