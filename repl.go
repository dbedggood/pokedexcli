package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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
	commands["map"] = cliCommand{
		name:        "map",
		description: "Display names of next 20 areas",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "map",
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
		url = "https://pokeapi.co/api/v2/location-area"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("error fetching data: %s", res.Status)
	}

	decodedStruct := struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
		} `json:"results"`
	}{}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedStruct); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range decodedStruct.Results {
		fmt.Println(area.Name)
	}

	nextUrl = decodedStruct.Next
	prevUrl = decodedStruct.Previous
	return nil
}
