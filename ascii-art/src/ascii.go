package ascii

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func GenerateAscii(s string) string {
	chars := LoadChars() //Load ascii characters from standard.txt
	lines := SplitByNewline(s)
	var outputRows [][][]string //3D slice to easily handle sections (i.e. "i'm section1 \n and i'm section 2" -> [<"i'm section 1 ">, <" and i'm section2">]), rows that make up sections (i.e. [<first row of section>, <second ...>, ..., <eigth row of section>]) and letters (i.e. [<top row of the letter a>, <top row of the letter b>, etc.])
	for _, line := range lines {
		outputRows = append(outputRows, ConvertToAscii([]byte(line), chars))
	}
	return RowsToText(outputRows)
}

func RowsToText(sections [][][]string) string {
	var output string
	for _, rows := range sections {
		for _, row := range rows {
			output += strings.Join(row, "") + "\n" //Joining each letter and adding newline to make a row, character by cahracter, row by row, section by section
		}
	}
	return output
}

func ConvertToAscii(text []byte, chars [][]string) [][]string {
	if len(text) < 1 { //Empty row handling
		return [][]string{{""}}
	}
	output := make([][]string, 8) //Each character has a height of 8 lines, thus exactly 8 row-slices are needed
	for _, b := range text {
		for i, line := range chars[b-32] { //b is the decimal ascii/byte code for a given letter, 32 is subtracted, since the first character in chars (space, index 0) has the decimal value of 32, and the rest are continuously sequential
			output[i] = append(output[i], line)
		}
	}
	return output
}

func LoadChars() [][]string {
	var output [][]string
	var char []string
	var active bool //Keeps track of whether the current line is part of a characters definition, this can be used instead of counting til every 8th row because no characters have gaps between the start and end of their definition

	lines := GetFileInput()

	for i, line := range lines {
		if line != "" && i+1 != len(lines) { //i+1 != len(lines) handles the adding of the last character in standard.txt, if it is not included and the last line of the file belongs to the last character, it will never be added
			active = true
			char = append(char, line)
		} else {
			if active {
				output = append(output, char)
				active = false
				char = []string{}
			}
		}
	}

	return output
}

func GetFileInput() []string {
	var output []string
	var readFile *os.File
	var err error

	if flag.Lookup("test.v") == nil { //If a "test.v" flag is present, the program was called via "go test" of "*_test.go", which can't be in the same folder as "main.go", and thus needs a different path to get to "standard.txt"
		readFile, err = os.Open("src/standard.txt")
	} else {
		readFile, err = os.Open("standard.txt")
	}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() { //Reading line by line
		output = append(output, scanner.Text())
	}

	readFile.Close()

	return output
}

func WriteToFile(filename string, text string) {
	err := os.WriteFile(filename, []byte(text), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func SplitByNewline(s string) []string {
	var lines []string
	var output []string
	lines = strings.Split(s, "\\n") //Works with cmd and bash calls but not "go test" calls
	for _, line := range lines {
		output = append(output, strings.Split(line, string(byte(10)))...) //Works with "go test" calls but not cmd and bash calls
	}
	return output
}
