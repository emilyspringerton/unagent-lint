package lint

import "unagent-lint/lint/result"

func Run(path string, cfg Config) result.Report {
	files := ScanFiles(path)
	report := result.NewReport()

	for _, file := range files {
		content := LoadFile(file)
		report.MarkFileChecked()

		for _, rule := range ActiveRules(cfg) {
			if v := rule.Check(content, file); v != nil {
				report.Add(*v)
			}
		}
	}
	if cfg.Strict && report.Warnings > 0 {
		report.Failures += report.Warnings
		report.Warnings = 0
	}
	return report
}
