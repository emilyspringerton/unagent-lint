package result

type Violation struct {
	Rule     string `json:"rule"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

type Report struct {
	Violations   []Violation
	Failures     int
	Warnings     int
	FilesChecked int
}

func NewReport() Report {
	return Report{}
}

func (r *Report) Add(v Violation) {
	r.Violations = append(r.Violations, v)
	if v.Severity == "FAIL" {
		r.Failures++
	}
	if v.Severity == "WARN" {
		r.Warnings++
	}
}

func (r *Report) MarkFileChecked() {
	r.FilesChecked++
}
