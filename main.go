package main

import (
	"fmt"
	"os"
)

func main() {
	input1 := os.Args[1]

	if input1 == "" {
		return
	}
	if input1 == "\\n" {
		fmt.Println()
		return
	}
	input2 := os.Args[2]
	Files_maped, err := Loadbanner(input2)

	if err != nil {
		fmt.Println("invalid")
	}
	render := Renderline(Files_maped)

	fmt.Print(render)
}
