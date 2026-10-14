
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := transform(string(content))

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Text transformed successfully!")
}

func transform(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = transformLine(strings.TrimSuffix(line, "\r"))
	}

	return strings.Join(lines, "\n")
}

func transformLine(line string) string {
	words := strings.Fields(line)
	result := []string{}

	for _, word := range words {
		marker := strings.TrimRight(word, ".,!?;:")

		if strings.HasPrefix(marker, "(") && strings.HasSuffix(marker, ")") {
			rule := strings.Trim(marker, "()")
			parts := strings.Split(rule, ",")
			action := strings.TrimSpace(parts[0])
			count := 1

			if len(parts) > 1 {
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil && n > 0 {
					count = n
				}
			}

			if action == "up" || action == "low" || action == "cap" {
				changed := 0

				for j := len(result) - 1; j >= 0 && changed < count; j-- {
					switch action {
					case "up":
						result[j] = strings.ToUpper(result[j])
					case "low":
						result[j] = strings.ToLower(result[j])
					case "cap":
						result[j] = capitalize(result[j])
					}
					changed++
				}
				continue
			}

			if action == "hex" || action == "bin" {
				if len(result) > 0 {
					last := len(result) - 1
					base := 16

					if action == "bin" {
						base = 2
					}

					number, err := strconv.ParseInt(result[last], base, 64)
					if err == nil {
						result[last] = strconv.FormatInt(number, 10)
					}
				}
				continue
			}
		}

		result = append(result, word)
	}

	text := strings.Join(result, " ")
	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixArticles(text)

	return text
}

func capitalize(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)

	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}

	return string(runes)
}

func fixPunctuation(text string) string {
	before := regexp.MustCompile(`\s+([.,!?;:]+)`)
	after := regexp.MustCompile(`([.,!?;:]+)(\S)`)

	text = before.ReplaceAllString(text, "$1")
	text = after.ReplaceAllString(text, "$1 $2")

	return text
}

func fixQuotes(text string) string {
	var result strings.Builder
	position := 0

	for position < len(text) {
		open := strings.Index(text[position:], "'")
		if open == -1 {
			result.WriteString(text[position:])
			break
		}

		open += position
		result.WriteString(text[position:open])
		result.WriteByte('\'')

		close := strings.Index(text[open+1:], "'")
		if close == -1 {
			result.WriteString(text[open+1:])
			break
		}

		close += open + 1
		inside := strings.TrimSpace(text[open+1 : close])
		result.WriteString(inside)
		result.WriteByte('\'')

		position = close + 1
	}

	return result.String()
}

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.EqualFold(words[i], "a") {
			next := strings.ToLower(words[i+1])

			if len(next) > 0 && strings.ContainsRune("aeiouh", rune(next[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
