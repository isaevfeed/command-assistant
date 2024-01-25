package helper

import "unicode"

func FirstLettertoUpper(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
