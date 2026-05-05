package main

import (
	"os"
	"strings"
)

func Renderline(font map[rune][]string) []string {
	result :=[]string{}
	input := strings.TrimSpace(os.Args[1])
	input1 := Split(input)

	for _, ch := range input1 {
		if len(result)>0 && ch=="" {
			result=append(result,"") 
			continue
		}
		for i := 0; i < 8; i++ {
			line := ""
			for _, char := range ch {
				line+= font[char][i]
			}
			result=append(result, line+"\n")
		}
	}
	return result
}

