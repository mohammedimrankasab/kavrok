package analyzer

// Severity represents the importance of a finding.
type Severity string

const (
	// SeverityCritical represents a critical cluster issue.
	SeverityCritical Severity = "critical"

	// SeverityWarning represents a potentially unhealthy condition.
	SeverityWarning Severity = "warning"

	// SeverityInfo represents informational feedback.
	SeverityInfo Severity = "info"
)

// Finding represents an engineering observation produced by Kavrok.
type Finding struct {
	Severity Severity
	Code     string
	Title    string
	Message  string
}
