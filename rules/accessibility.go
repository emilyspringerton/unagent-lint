package rules

import "strings"

func AccessibilityFraming() Rule {
	return Rule{
		ID:       "A11Y-001",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			required := []string{"accessibility", "assistive", "screen reader", "inclusive", "disability"}
			blocked := []string{"productivity", "optimize", "automation", "efficiency"}

			if _, found := containsAny(lowered, blocked); found {
				return NewViolation("A11Y-001", "FAIL",
					"UNAGENT must be framed as accessibility software, not productivity tooling.",
					file, 0)
			}

			if _, found := containsAny(lowered, required); !found {
				return NewViolation("A11Y-001", "FAIL",
					"UNAGENT must be framed as accessibility software, not productivity tooling.",
					file, 0)
			}
			return nil
		},
	}
}

func NonDeficitLanguage() Rule {
	return Rule{
		ID:       "A11Y-002",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"broken", "fix you", "deficit", "sufferer", "incapable"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("A11Y-002", "FAIL",
					"Deficit-based language detected. Accessibility tools support users; they do not fix them.",
					file, line)
			}
			return nil
		},
	}
}
