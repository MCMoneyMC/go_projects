package src

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CorrectText(input string) string {
	textArr := regexp.MustCompile(`\([^)]*\)|\S+`).FindAllString(input, -1)
	textArr = ConvertTags(textArr)
	text := strings.Join(textArr, " ")
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixArticles(text)
	if string(text[len(text)-1]) == " " {
		text = text[:len(text)-1]
	}
	return text
}

func GetInput(filename string) string {
	text, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}

func ConvertTags(text []string) []string {
	var output []string
	for _, word := range text {
		output = Build(word, output)
	}
	return output
}

func Build(s string, output []string) []string {
	switch s {
	case "(hex)":
		if len(output) > 0 {
			hex, err := strconv.ParseInt((output[len(output)-1]), 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			output[len(output)-1] = strconv.Itoa(int(hex))
		}
		return output
	case "(bin)":
		if len(output) > 0 {
			bin, err := strconv.ParseInt((output[len(output)-1]), 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			output[len(output)-1] = strconv.Itoa(int(bin))
		}
		return output
	case "(up)":
		if len(output) > 0 {
			output[len(output)-1] = strings.ToUpper(output[len(output)-1])
		}
		return output
	case "(low)":
		if len(output) > 0 {
			output[len(output)-1] = strings.ToLower(output[len(output)-1])
		}
		return output
	case "(cap)":
		if len(output) > 0 {
			output[len(output)-1] = Capitalize(output[len(output)-1])
		}
		return output
	}
	upSpecial, _ := regexp.Match(`\(up, [0-9]*\)`, []byte(s))
	if upSpecial {
		count, err := strconv.Atoi(s[5 : len(s)-1])

		if err != nil {
			log.Fatal(err)
		}

		if len(output) > 0 {
			for i := 0; i < count; i++ {
				output[len(output)-1-i] = strings.ToUpper(output[len(output)-1-i])
			}
		}
		return output
	}
	lowSpecial, _ := regexp.Match(`\(low, [0-9]*\)`, []byte(s))
	if lowSpecial {
		count, err := strconv.Atoi(s[6 : len(s)-1])

		if err != nil {
			log.Fatal(err)
		}

		if len(output) > 0 {
			for i := 0; i < count; i++ {
				output[len(output)-1-i] = strings.ToLower(output[len(output)-1-i])
			}
		}
		return output
	}
	capSpecial, _ := regexp.Match(`\(cap, [0-9]*\)`, []byte(s))
	if capSpecial {
		count, err := strconv.Atoi(s[6 : len(s)-1])

		if err != nil {
			log.Fatal(err)
		}

		if len(output) > 0 {
			for i := 0; i < count; i++ {
				output[len(output)-1-i] = Capitalize(output[len(output)-1-i])
			}
		}
		return output
	}
	output = append(output, s)
	return output
}

func Capitalize(s string) string {
	var output string
	if 'a' <= rune(s[0]) && rune(s[0]) <= 'z' {
		output += string(rune(s[0]) + 'A' - 'a')
	}
	for _, r := range s[1:] {
		if 'A' <= r && r <= 'Z' {
			output += string(r + 'a' - 'A')
		} else {
			output += string(r)
		}
	}
	return output
}

func FixPunctuation(s string) string {
	punctuationRegex := regexp.MustCompile(`\s*(\.\.\.|!\?|[.,;:!?])\s*`)
	correctedText := punctuationRegex.ReplaceAllString(s, "$1 ")
	return correctedText
}

func FixQuotes(s string) string {
	re := regexp.MustCompile(`'([^']+)'`)
	s = re.ReplaceAllStringFunc(s, RemoveSpaces)

	return s
}

func RemoveSpaces(s string) string {
	return "'" + strings.Trim(s[1:len(s)-2], " ") + "'"
}

func FixArticles(s string) string {
	indices := regexp.MustCompile(` (a|A) [aieouhAIEUOH]`).FindAllIndex([]byte(s), -1)
	for n, i := range indices {
		if rune(s[i[0]+1+n]) == 'a' {
			s = s[:i[0]+1+n] + "an" + s[i[0]+2+n:]
		} else {
			s = s[:i[0]+1+n] + "An" + s[i[0]+2+n:]
		}
	}
	return s
}

func WriteToFile(s string, filename string) {
	err := os.WriteFile(filename, []byte(s), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
