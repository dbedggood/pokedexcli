package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

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
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect a Pokemon",
		callback:    commandInspect,
	}
	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Display information about a Pokemon",
		callback:    commandPokedex,
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

	fmt.Println("Exploring " + locationAreaName + "...")

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

var pokedex map[string]Pokemon

func commandCatch(args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]

	pokemonSpeciesUrl := POKEMON_SPECIES_BASE_URL + pokemonName
	pokemonSpecies := PokemonSpecies{}
	if err := pokeapi.Fetch(pokemonSpeciesUrl, &pokemonSpecies); err != nil {
		return err
	}

	pokemonUrl := POKEMON_BASE_URL + pokemonName
	pokemon := Pokemon{}
	if err := pokeapi.Fetch(pokemonUrl, &pokemon); err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")

	time.Sleep(1 * time.Second)

	if wasPokemonCaught(pokemonSpecies.CaptureRate) {
		fmt.Println(pokemon.Name + " was caught!")

		if pokedex == nil {
			pokedex = map[string]Pokemon{}
		}
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}

	return nil
}

func wasPokemonCaught(captureRate int) bool {
	r := rand.Float64()
	return r < float64(captureRate)/255.0
}

func commandInspect(args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := pokedex[pokemonName]
	if !ok {
		return errors.New("pokemon not caught yet")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, type_ := range pokemon.Types {
		fmt.Printf(" - %s\n", type_.Type.Name)
	}

	return nil
}

func commandPokedex(args []string) error {
	if len(args) > 0 {
		return errors.New("don't give me an arg")
	}

	if len(pokedex) == 0 {
		return errors.New("you haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range pokedex {
		fmt.Println(" - " + name)
	}

	return nil
}
