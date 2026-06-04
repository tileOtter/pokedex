package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tileOtter/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	fields := strings.Fields(lowered)
	return fields
}

func startRepl(cfg *pokeapi.Config) {
	cmdMap := getCommandsMap()
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %s", err)
	}
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
