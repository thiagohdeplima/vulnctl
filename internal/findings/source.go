package findings

import (
	"context"
	"fmt"

	"github.com/thiagohdeplima/riskctl/internal/config"
)

// FindingSource retrieves vulnerability findings from an external provider.
type FindingSource interface {
	ListFindings(ctx context.Context, cfg config.Config, query Query) ([]Finding, error)
}

// StubFindingSource is a stub that always returns not implemented.
type StubFindingSource struct{}

func (s *StubFindingSource) ListFindings(_ context.Context, _ config.Config, _ Query) ([]Finding, error) {
	return nil, fmt.Errorf("not implemented")
}
