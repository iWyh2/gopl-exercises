package ch3

import (
	"bytes"
	"strings"
)

func CommaModified(str string) string {
	symbol := str[0]
	var buf bytes.Buffer
	buf.WriteString(string(symbol))
	indexOfDot := strings.LastIndex(str, ".")
	str1 := Comma(str[1:indexOfDot])
	buf.WriteString(str1)
	buf.WriteString(".")
	str2 := Comma(str[indexOfDot+1:])
	buf.WriteString(str2)
	return buf.String()
}
