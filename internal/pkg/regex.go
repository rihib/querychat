package pkg

import (
	"fmt"
	"regexp"
)

/*
FindPattern is a function that finds a pattern in the text, and returns the first match.
*/
func FindPattern(text, pattern string) (string, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	matches := regex.FindStringSubmatch(text)
	if len(matches) <= 1 {
		return "", fmt.Errorf("no match found, text: %s, pattern: %s", text, pattern)
	}
	return matches[1], nil
}
