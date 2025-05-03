package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	words := strings.Fields(cleaned)
	return words
}
