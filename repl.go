package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/dbedggood/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(args []string) error {
			return commandHelp(args, commands)
		},
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
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a new area",
		callback:    commandExplore,
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

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if err := command.callback(args); err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	words := strings.Fields(lowercaseText)
	return words
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args []string, commands map[string]cliCommand) error {
	if len(args) > 0 {
		fmt.Println("Usage: help")
		return nil
	}

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

const LOCATION_AREA_BASE_URL = "https://pokeapi.co/api/v2/location-area/"

var nextUrl string
var prevUrl string

func commandMap(args []string) error {
	if len(args) > 0 {
		fmt.Println("Usage: map")
		return nil
	}

	if err := fetchAndDisplayAreas(nextUrl); err != nil {
		return err
	}
	return nil
}

func commandMapBack(args []string) error {
	if len(args) > 0 {
		fmt.Println("Usage: mapb")
		return nil
	}

	if err := fetchAndDisplayAreas(prevUrl); err != nil {
		return err
	}
	return nil
}

func fetchAndDisplayAreas(url string) error {
	if url == "" {
		url = LOCATION_AREA_BASE_URL + "?offset=0&limit=20"
	}

	locationAreas := LocationAreas{}
	if err := pokeapi.Fetch(url, &locationAreas); err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	nextUrl = locationAreas.Next
	prevUrl = locationAreas.Previous
	return nil
}

func commandExplore(args []string) error {
	if len(args) == 0 {
		return errors.New("argument cannot be empty")
	}
	locationAreaName := args[0]

	fmt.Println("Exploring pastoria-city-area...")

	locationAreaDetails := LocationAreaDetails{}
	if err := pokeapi.Fetch(LOCATION_AREA_BASE_URL+locationAreaName, &locationAreaDetails); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationAreaDetails.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}

	return nil
}
