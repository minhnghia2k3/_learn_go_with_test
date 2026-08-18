package iterations

import "strings"

// Repeat the character in count times.
func Repeat(char string, count int) string {
	var repeated strings.Builder

	for i := 0; i < count; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}
