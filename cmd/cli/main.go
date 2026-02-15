package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thiagohdeplima/riskctl/cmd/cli/findings"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "riskctl",
	Short: "Multi-source vulnerability aggregation CLI",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(findings.FindingCmd)
}
