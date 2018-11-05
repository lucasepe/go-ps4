package utils

import (
	"strings"
	"unicode/utf8"
)

func times(str string, n int) (out string) {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(str)
	}
	return sb.String()
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func LeftPad(str string, len int, pad string) string {
	return times(pad, len-utf8.RuneCountInString(str)) + str
}

// Right right-pads the string with pad up to len runes
func RightPad(str string, len int, pad string) string {
	return str + times(pad, len-utf8.RuneCountInString(str))
}
