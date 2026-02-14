package findings

import "github.com/spf13/cobra"

// Cmd is the root command for managing vulnerability findings.
var Cmd = &cobra.Command{
	Use:   "findings",
	Short: "Manage vulnerability findings",
}
