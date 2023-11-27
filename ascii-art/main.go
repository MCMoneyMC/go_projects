package main

import (
	ascii "ascii-art/src"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		for _, arg := range args {
			fmt.Println(arg)
		}
		panic("This program requires at least one, and accepts up to 2 arguments")
	}

	text := ascii.GenerateAscii(args[0])
	if len(args) == 1 {
		fmt.Print(text)
	} else {
		ascii.WriteToFile("output.txt", text)
	}
}
