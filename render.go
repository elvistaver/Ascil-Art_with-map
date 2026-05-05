package main

import (
	"os"
	"strings"
)

func Renderline(font map[rune][]string) string {
	result := ""

	input := strings.TrimSpace(os.Args[1])
	input1 := Split(input)


	for _, ch := range input1 {
		if ch =="" {
			result+="\n" 
			continue
		}
		for i := 0; i < 8; i++ {
			line := ""
			for _, char := range ch {
				line+= font[char][i]
			}
			result +=line+"\n"
		}
	}
	return result
}

