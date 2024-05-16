package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string = ""
	Commit  string = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\ncommit: %s\n", Version, Commit)
	},
}
