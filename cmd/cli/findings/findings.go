package findings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/thiagohdeplima/riskctl/internal/config"
	findingspkg "github.com/thiagohdeplima/riskctl/internal/findings"
	"github.com/thiagohdeplima/riskctl/internal/findings/sources"
)

var FindingCmd = &cobra.Command{
	Use:   "findings",
	Short: "Manage vulnerability findings",
}

var FindingListCmd = &cobra.Command{
	Use:   "list",
	Short: "retrieve vulnerabilities from a source",
	Run:   ExecFindingListCmd,
}

func init() {
	FindingCmd.AddCommand(FindingListCmd)
}

func ExecFindingListCmd(cmd *cobra.Command, args []string) {
	loader := config.ConfigLoader{}
	cfg, err := loader.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	src := sources.NewAWSInspectorSource()
	results, err := src.ListFindings(cmd.Context(), cfg, findingspkg.Query{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Found %d findings\n\n", len(results))
	for _, f := range results {
		fmt.Printf("%-14s %s\n", "FindingID:", f.FindingID)
		fmt.Printf("%-14s %s\n", "Severity:", f.Severity)
		fmt.Printf("%-14s %s\n", "VulnID:", f.VulnID)
		fmt.Printf("%-14s %s\n", "AssetType:", f.AssetType)
		fmt.Printf("%-14s %s\n", "AssetID:", f.AssetID)
		fmt.Printf("%-14s %v\n", "FixAvailable:", f.FixAvailable)
		fmt.Println("---")
	}
}
