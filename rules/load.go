package rules

import "strings"

func CognitiveLoadReduction() Rule {
	return Rule{
		ID:       "LOAD-001",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"choose one of", "pick any", "multiple options"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("LOAD-001", "FAIL",
					"Content increases cognitive load instead of reducing it.",
					file, line)
			}
			return nil
		},
	}
}

func DefaultsFirst() Rule {
	return Rule{
		ID:       "LOAD-002",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"configure first", "set up before", "must configure"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("LOAD-002", "FAIL",
					"Users must configure before receiving relief. Defaults are required.",
					file, line)
			}
			return nil
		},
	}
}

func OfflineSafety() Rule {
	return Rule{
		ID:       "LOAD-003",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"asap", "urgent", "immediately"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("LOAD-003", "FAIL",
					"Content introduces urgency or guilt around availability.",
					file, line)
			}
			return nil
		},
	}
}
