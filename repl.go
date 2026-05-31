package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	trimmed := []string{}
	for _, s := range split {
		trimmed = append(trimmed, strings.TrimSpace(s))
	}
	return split
}
