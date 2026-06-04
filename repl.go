package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tileOtter/pokedex/internal/pokeapi"
)

type commandsList struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config) error
}

func getCommandsMap() map[string]commandsList {
	commands := map[string]commandsList{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 map locations",
			callback:    commandMapb,
		},
	}
	return commands
}

func commandExit(cfg *pokeapi.Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeapi.Config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	cmdMap := getCommandsMap()
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *pokeapi.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	decoded, err := pokeapi.FetchLocations(url)
	if err != nil {
		log.Fatalf("error fetching locations: %s", err)
	}
	cfg.Next = decoded.Next
	cfg.Previous = decoded.Previous
	for _, result := range decoded.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return nil
}

func commandMapb(cfg *pokeapi.Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.Previous

	decoded, err := pokeapi.FetchLocations(url)
	if err != nil {
		log.Fatalf("error fetching locations: %s", err)
	}
	cfg.Next = decoded.Next
	cfg.Previous = decoded.Previous
	for _, result := range decoded.Results {
		fmt.Printf("%s\n", result.Name)
	}
	return nil
}

func startRepl(cfg *pokeapi.Config) {
	cmdMap := getCommandsMap()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedStringMap := cleanInput(input)
		commandInput := cleanedStringMap[0]
		command, ok := cmdMap[commandInput]
		if ok {
			command.callback(cfg)
		} else {
			fmt.Print("Unknown command\n")
		}

	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}
