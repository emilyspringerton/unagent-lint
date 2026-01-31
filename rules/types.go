package rules

import "unagent-lint/lint/result"

type Violation = result.Violation

type Rule struct {
	ID       string
	Severity string
	Check    func(content, file string) *Violation
}

func NewViolation(ruleID, severity, message, file string, line int) *Violation {
	return &Violation{
		Rule:     ruleID,
		Severity: severity,
		Message:  message,
		File:     file,
		Line:     line,
	}
}
