package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tileOtter/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
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
