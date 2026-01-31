package rules

import "strings"

func NoDeception() Rule {
	return Rule{
		ID:       "TRUST-001",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"we already", "we have handled", "we took care of"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("TRUST-001", "FAIL",
					"Content implies commitments without user consent.",
					file, line)
			}
			return nil
		},
	}
}

func ConsentBoundary() Rule {
	return Rule{
		ID:       "TRUST-002",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"we will proceed", "we will act", "taking action"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("TRUST-002", "FAIL",
					"Action described without explicit user consent.",
					file, line)
			}
			return nil
		},
	}
}
