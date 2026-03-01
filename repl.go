package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mandarinomczumo/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	cleanText := []string{}
	for word := range strings.FieldsSeq(text) {
		cleanText = append(cleanText, strings.ToLower(word))
	}
	return cleanText
}

type config struct {
	pokeapiClient        pokeapi.Client
	LocationAreaNext     string
	LocationAreaPrevious string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	exitCmd := cliCommand{name: "exit", description: "Exit the Pokedex", callback: commandExit}
	helpCmd := cliCommand{name: "help", description: "Displays a help message", callback: commandHelp}
	mapCmd := cliCommand{name: "map", description: "Displays 20 location areas in the Pokemon world", callback: commandMap}
	mapBackCmd := cliCommand{name: "mapb", description: "Displays the next 20 location areas in the Pokemon world", callback: commandMapBack}
	return map[string]cliCommand{"exit": exitCmd, "help": helpCmd, "map": mapCmd, "mapb": mapBackCmd}
}

func startRepl(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		input := cleanInput(userInput)[0]

		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println("error", err)
			continue
		}
	}
}
