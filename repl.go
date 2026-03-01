package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanText := []string{}
	for word := range strings.FieldsSeq(text) {
		cleanText = append(cleanText, strings.ToLower(word))
	}
	return cleanText
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	exitCmd := cliCommand{name: "exit", description: "Exit the Pokedex", callback: commandExit}
	helpCmd := cliCommand{name: "help", description: "Displays a help message", callback: commandHelp}
	return map[string]cliCommand{"exit": exitCmd, "help": helpCmd}
}

func main() {
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
		err := command.callback()
		if err != nil {
			fmt.Println("error", err)
			continue
		}
	}
}
