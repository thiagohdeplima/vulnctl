package findings

import (
	"fmt"

	"github.com/thiagohdeplima/riskctl/internal/config"
)

// FindingEngine processes findings by applying policies, exclusions, and enrichment.
type FindingEngine interface {
	Process(cfg config.Config, findings []Finding, query Query) ([]Finding, error)
}

// StubFindingEngine is a stub that always returns not implemented.
type StubFindingEngine struct{}

func (s *StubFindingEngine) Process(_ config.Config, _ []Finding, _ Query) ([]Finding, error) {
	return nil, fmt.Errorf("not implemented")
}
