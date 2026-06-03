package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commandsList struct {
	name        string
	description string
	callback    func() error
}

func getCommandsList() map[string]commandsList {
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
	}
	return commands
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	cmdList := getCommandsList()
	for _, cmd := range cmdList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedStringMap := cleanInput(input)
		commandInput := cleanedStringMap[0]
		commands := getCommandsList()
		command, ok := commands[commandInput]
		if ok {
			command.callback()
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
