package main

import "strings"

func Split(text string) []string {
	newline := strings.Split(text, "\\n")
	return newline
}
