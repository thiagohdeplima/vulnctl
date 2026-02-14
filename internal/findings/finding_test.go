package findings_test

import (
	"testing"
	"time"

	"github.com/thiagohdeplima/riskctl/internal/findings"
)

func TestFindingInstantiation(t *testing.T) {
	now := time.Now()

	f := findings.Finding{
		FindingID:    "FINDING-001",
		Source:       "aws-inspector",
		AssetType:    "ec2",
		AssetID:      "i-0123456789abcdef0",
		VulnID:       "CVE-2024-1234",
		Severity:     "HIGH",
		FixAvailable: true,
		FirstSeen:    now,
		LastSeen:     now,
		Excluded:     false,
		ExcludedBy:   "",
	}

	if f.FindingID != "FINDING-001" {
		t.Errorf("expected FindingID FINDING-001, got %s", f.FindingID)
	}

	if f.Source != "aws-inspector" {
		t.Errorf("expected Source aws-inspector, got %s", f.Source)
	}
}
