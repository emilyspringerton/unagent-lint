package rules

import "strings"

func findLine(content, phrase string) int {
	index := strings.Index(strings.ToLower(content), strings.ToLower(phrase))
	if index == -1 {
		return 0
	}
	return 1 + strings.Count(content[:index], "\n")
}

func containsAny(lowered string, phrases []string) (string, bool) {
	for _, phrase := range phrases {
		if strings.Contains(lowered, phrase) {
			return phrase, true
		}
	}
	return "", false
}
