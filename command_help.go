package main

import (
	"fmt"
)

func commandHelp(c *config, args ...string) error {
	commands := getCommands()
	fmt.Println(`Welcome to the Pokedex!
Usage:`)
	for _, cmd := range commands {
		fmt.Println(cmd.name+": ", cmd.description)
	}
	return nil
}
