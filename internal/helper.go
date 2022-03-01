package internal

import (
	"strings"
)

func isOnlyContain(s, v string) bool {
	return strings.Trim(s, v) == ""
}

func isValidDirName(s string) bool {
	if isOnlyContain(s, ".") {
		return false
	}

	return true
}

func inText(key, text string) bool {
	matchCount := 0
	keyIndex := 0
	textIndex := 0
	for keyIndex < len(key) && keyIndex < len(text) && textIndex < len(text) {
		keyRune := key[keyIndex]
		textRune := text[textIndex]
		if keyRune == textRune {
			matchCount++
			keyIndex++
		}
		textIndex++
	}

	return len(key) == matchCount
}

func splitSpace(s string) []string {
	s = strings.Trim(s, " ")
	parmChars := []rune(s)
	inSingleQuote := false
	inDoubleQuote := false
	for index, ch := range parmChars {
		if ch == '"' && !inSingleQuote {
			inDoubleQuote = !inDoubleQuote
			parmChars[index] = '\n'
		}
		if parmChars[index] == '\'' && !inDoubleQuote {
			inSingleQuote = !inSingleQuote
			parmChars[index] = '\n'
		}
		if !inSingleQuote && !inDoubleQuote && parmChars[index] == ' ' {
			parmChars[index] = '\n'
		}
	}
	list := strings.Split(string(parmChars), "\n")
	var result []string
	for _, s := range list {
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return result
}
