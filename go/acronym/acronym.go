package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate : accepts a string and returns the acronym.
func Abbreviate(s string) string {
	var result = ""
	re := regexp.MustCompile("\\s|-")
	words := re.Split(s, -1)

	for _, v := range words {
		result += v[:1]
	}
	return strings.ToUpper(result)
}
