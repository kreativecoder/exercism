package bob

import (
	"strings"
	"unicode"
)

//Hey : returns response based on a set of conditions
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	} else if strings.ToUpper(remark) == remark && HasLetter(remark) && strings.HasSuffix(remark, "?") {
		return "Calm down, I know what I'm doing!"
	} else if strings.ToUpper(remark) == remark && HasLetter(remark) {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	return "Whatever."
}

// HasLetter : checks if a string has at least one letter
func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
