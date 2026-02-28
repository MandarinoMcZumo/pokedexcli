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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// func commandHelp() {
// 	exitCmd := cliMedatada{name: "exit", description: "Exit the Pokedex"}
// 	helpCmd := cliMedatada{name: "help", description: "Displays a help message"}
// 	commands := []cliMedatada{exitCmd, helpCmd}
// 	fmt.Println(`Welcome to the Pokedex!
// 			Usage:`)
// 	for _, cmd := range commands {
// 		fmt.Println(cmd.name+": ", cmd.description)
// 	}
// }

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{}
	exitCmd := cliCommand{name: "exit", description: "Exit the Pokedex", callback: commandExit}
	helpCmd := cliCommand{name: "help", description: "Displays a help message", callback: func() error {
		fmt.Println(`Welcome to the Pokedex!
Usage:`)
		for _, cmd := range commands {
			fmt.Println(cmd.name+": ", cmd.description)
		}
		return nil
	}}
	commands["exit"] = exitCmd
	commands["help"] = helpCmd
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		input := cleanInput(userInput)[0]

		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
		}
		err := command.callback()
		if err != nil {
			fmt.Println("error", err)
		}
	}
}
