package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2023")

	funcs := map[string]func(filepath string){
		"18": day18,
	}

	args := os.Args[1:]

	if len(args) < 1 {
		panic("Provide a day number")
	}
	day := args[0]
	fmt.Printf("Day %s\n", day)

	filepath := fmt.Sprintf("example/day%s.txt", day)
	if len(args) > 1 {
		filepath = args[1]
	}
	fmt.Printf("File: %s\n\n", filepath)

	funcs[day](filepath)
}
