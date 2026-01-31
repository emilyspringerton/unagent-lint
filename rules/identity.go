package rules

import "strings"

func NotAnAgent() Rule {
	return Rule{
		ID:       "ID-001",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			blocked := []string{
				"acts on your behalf",
				"autonomous agent",
				"handles everything",
				"makes decisions for you",
			}
			lowered := strings.ToLower(content)
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("ID-001", "FAIL",
					"UNAGENT must not be described as an autonomous agent.",
					file, line)
			}
			return nil
		},
	}
}

func NoImpersonation() Rule {
	return Rule{
		ID:       "ID-002",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			blocked := []string{
				"impersonate",
				"pretend to be you",
				"speak as you",
				"on your behalf",
			}
			lowered := strings.ToLower(content)
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("ID-002", "FAIL",
					"UNAGENT may not impersonate or replace the user.",
					file, line)
			}
			return nil
		},
	}
}
