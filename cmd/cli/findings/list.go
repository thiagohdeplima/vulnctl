package findings

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/riskctl/internal/config"
	findingspkg "github.com/thiagohdeplima/riskctl/internal/findings"
	"github.com/thiagohdeplima/riskctl/internal/renderer"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List vulnerability findings",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Instantiate stub implementations to prove the wiring compiles.
		var _ config.ConfigLoader = &config.StubConfigLoader{}
		var _ findingspkg.FindingSource = &findingspkg.StubFindingSource{}
		var _ findingspkg.FindingEngine = &findingspkg.StubFindingEngine{}
		var _ renderer.Renderer = &renderer.StubRenderer{}

		fmt.Println("not implemented")
		return nil
	},
}

func init() {
	Cmd.AddCommand(listCmd)
}
