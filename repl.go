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
	pokedex              map[string]pokeapi.Pokemon
	LocationAreaNext     string
	LocationAreaPrevious string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	exitCmd := cliCommand{name: "exit", description: "Exit the Pokedex", callback: commandExit}
	helpCmd := cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	mapCmd := cliCommand{
		name:        "map",
		description: "Displays 20 location areas in the Pokemon world",
		callback:    commandMap,
	}
	mapBackCmd := cliCommand{
		name:        "mapb",
		description: "Displays the next 20 location areas in the Pokemon world",
		callback:    commandMapBack,
	}
	exploreCmd := cliCommand{
		name:        "explore",
		description: "Explores a region",
		callback:    commandExplore,
	}
	catchCommand := cliCommand{
		name:        "catch",
		description: "Throw pokeball to pokemon",
		callback:    commandCatch,
	}
	inspectCommand := cliCommand{
		name:        "inspect",
		description: "Inspect pokemon (Must have been captured previously)",
		callback:    commandInspect,
	}
	return map[string]cliCommand{
		"exit":    exitCmd,
		"help":    helpCmd,
		"map":     mapCmd,
		"mapb":    mapBackCmd,
		"explore": exploreCmd,
		"catch":   catchCommand,
		"inspect": inspectCommand,
	}
}

func startRepl(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		words := cleanInput(userInput)

		command, ok := commands[words[0]]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
