package ch3

import (
	"bytes"
)

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	if n%3 != 0 {
		i := n % 3
		j := 0
		for ; i <= len(s); i += 3 {
			buf.WriteString(s[j:i])
			buf.WriteString(",")
			j = i
		}
	} else {
		n := 3
		j := 0
		for ; n <= len(s); n += 3 {
			buf.WriteString(s[j:n])
			buf.WriteString(",")
			j = n
		}
	}
	return buf.String()[:len(buf.String())-1]
}
