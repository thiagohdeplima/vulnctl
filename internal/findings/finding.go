package findings

import "time"

// Finding is the canonical vulnerability finding representation.
type Finding struct {
	FindingID    string
	Source       string
	AssetType    string
	AssetID      string
	VulnID       string
	Severity     string
	FixAvailable bool
	FirstSeen    time.Time
	LastSeen     time.Time
	Excluded     bool
	ExcludedBy   string
}

// Query defines filters for retrieving findings.
type Query struct {
	Source   string
	Severity string
	Since    *time.Time
}
