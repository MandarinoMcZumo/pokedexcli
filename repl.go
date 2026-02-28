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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		clean := cleanInput(userInput)
		fmt.Println("Your command was: " + clean[0])
	}
}
