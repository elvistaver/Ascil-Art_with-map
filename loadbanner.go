package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func Loadbanner(filename string) (map[rune][]string, error) {
	lines := []string{}

	file, err := os.Open("bannerfiles/" + filename + ".txt")

	if err != nil {
		log.Fatal("\nfile not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err != scanner.Err() {
		fmt.Println("invalid file")
	}
	font:=make(map[rune][]string)
	for i:=0; i<len(lines); i+=9{
			Block:=lines[i:i+9]
			char:=rune((i/9)+32)
			font[char]=Block	
	}
	return font, nil
}
