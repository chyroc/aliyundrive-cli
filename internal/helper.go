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
	res := []string{}
	buf := []rune{}
	z := false
	for _, v := range []rune(s) {
		if v == '\\' {
			z = true
			continue
		} else if v == ' ' {
			if z {
				buf = append(buf, v)
				z = false
			} else {
				z = false
				if len(buf) > 0 {
					res = append(res, string(buf))
					buf = []rune{}
				}
			}
		} else {
			buf = append(buf, v)
		}
	}
	if len(buf) > 0 {
		res = append(res, string(buf))
	}
	return res
}
