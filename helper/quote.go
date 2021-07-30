package helper

import (
	"bytes"
)

// Quote wrap data with backticks`
func Quote(data string) string {
	return "`" + escapeBackticks(data) + "`"
}

func escapeBackticks(data string) string {
	var buf bytes.Buffer
	for _, c := range data {
		if c == '`' {
			buf.WriteByte('`')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
