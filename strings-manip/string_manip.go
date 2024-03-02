package stringsmanip

import (
	"strings"
)

// Compare two strings and returns int (0 if equal) || (-1 if a < b) || (1 if a > b)
func CompareString(string1, string2 string) int {
	return (strings.Compare(string1, string2))
}
