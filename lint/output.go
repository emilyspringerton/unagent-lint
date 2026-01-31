package lint

import (
	"encoding/json"
	"fmt"

	"unagent-lint/lint/result"
)

type jsonSummary struct {
	FilesChecked int `json:"files_checked"`
	Failures     int `json:"failures"`
	Warnings     int `json:"warnings"`
}

type jsonReport struct {
	Summary    jsonSummary        `json:"summary"`
	Violations []result.Violation `json:"violations"`
}

func Print(report result.Report, cfg Config) {
	if cfg.Format == "json" {
		payload := jsonReport{
			Summary: jsonSummary{
				FilesChecked: report.FilesChecked,
				Failures:     report.Failures,
				Warnings:     report.Warnings,
			},
			Violations: report.Violations,
		}
		encoded, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			fmt.Println("{}")
			return
		}
		fmt.Println(string(encoded))
		return
	}

	if len(report.Violations) == 0 {
		if !cfg.Quiet {
			fmt.Println("PASS: no violations found")
		}
		return
	}

	for _, v := range report.Violations {
		lineInfo := ""
		if v.Line > 0 {
			lineInfo = fmt.Sprintf(":%d", v.Line)
		}
		fmt.Printf("%s %s %s%s - %s\n", v.Severity, v.Rule, v.File, lineInfo, v.Message)
	}
}
