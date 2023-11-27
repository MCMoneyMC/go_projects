package main

import (
	ascii "ascii-art-fs/src"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		args = append(args, "standard")
	}

	if len(args) == 2 && (args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		text := ascii.GenerateAscii(args[0], args[1])
		fmt.Print(text)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
}
