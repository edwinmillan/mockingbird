package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mock(text string) string {
	textHolder := make([]string, 0)
	for i, letter := range strings.ToUpper(text) {
		letter := string(letter)
		if i%2 == 0 {
			letter = strings.ToLower(letter)
		}
		textHolder = append(textHolder, letter)
	}
	return strings.Join(textHolder, "")
}

func input(question string) string {
	textInput := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	text, _ := textInput.ReadString('\n')
	return text
}

func main() {
	fmt.Println("Enjoy Mocking! Enter Q to quit.")
	for {
		textInput := input("Enter text to Mock: ")
		if strings.TrimSpace(textInput) == strings.ToLower("q") {
			break
		}
		fmt.Println()
		fmt.Println(mock(textInput))
	}
}
