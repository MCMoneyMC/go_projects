package main

import (
	reloaded "go-reloaded/src"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("This program requires exactly 2 parameters--the names of the input and output file.")
	}
	input := reloaded.GetInput(args[0])
	text := reloaded.CorrectText(input)
	reloaded.WriteToFile(text, args[1])
}
