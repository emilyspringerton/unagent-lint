package rules

import (
	"strings"
)

func CalmAuthority() Rule {
	return Rule{
		ID:       "VOICE-001",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"sorry", "apologize", "i'm just", "trust me"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("VOICE-001", "FAIL",
					"Emily Voice requires calm authority without apology or dominance.",
					file, line)
			}
			return nil
		},
	}
}

func EmotionalTemperature() Rule {
	return Rule{
		ID:       "VOICE-002",
		Severity: "FAIL",
		Check: func(content, file string) *Violation {
			lowered := strings.ToLower(content)
			blocked := []string{"!!!", "???", "obviously", "ridiculous"}
			if phrase, found := containsAny(lowered, blocked); found {
				line := findLine(content, phrase)
				return NewViolation("VOICE-002", "FAIL",
					"Output escalates emotional temperature. Emily Voice de-escalates by default.",
					file, line)
			}
			return nil
		},
	}
}

func PrecisionOverVolume() Rule {
	return Rule{
		ID:       "VOICE-003",
		Severity: "WARN",
		Check: func(content, file string) *Violation {
			words := strings.Fields(content)
			if len(words) > 300 {
				return NewViolation("VOICE-003", "WARN",
					"Output may be longer than necessary. Consider reducing volume.",
					file, 0)
			}
			return nil
		},
	}
}
