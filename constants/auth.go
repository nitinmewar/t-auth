package constants

var Headers = struct {
	OrgHeader     string
	OrgAccHeader  string
	AUTHORIZATION string
	AccessSecret  string
	Sandbox       string
	TraceID       string
}{
	OrgHeader:     "x-studymitr-org",
	OrgAccHeader:  "x-studymitr-org-acc",
	AUTHORIZATION: "Authorization",
	AccessSecret:  "x-studymitr-key",
	Sandbox:       "sandbox",
	TraceID:       "x-studymitr-trace-id",
}
