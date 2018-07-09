package twofer

import (
	"strings"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
