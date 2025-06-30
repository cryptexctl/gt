package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "Command-line tool for Gitea",
	Long:  "gt is a CLI for interacting with a Gitea server: repositories, issues, PRs and more.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("host", "H", "", "Gitea host URL")
	rootCmd.PersistentFlags().StringP("token", "t", "", "API token")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
}
