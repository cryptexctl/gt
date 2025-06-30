package cmd

import (
	"fmt"

	"github.com/cryptexctl/teah/internal/config"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication",
	Long:  "Authentication commands for Gitea.",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Set authentication",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			return err
		}

		host, _ := cmd.Flags().GetString("host")
		token, _ := cmd.Flags().GetString("token")

		if host == "" {
			return fmt.Errorf("host is required (--host)")
		}
		if token == "" {
			return fmt.Errorf("token is required (--token)")
		}

		if err := config.Set("host", host); err != nil {
			return err
		}
		if err := config.Set("token", token); err != nil {
			return err
		}

		fmt.Printf("Authentication saved for %s\n", host)
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			return err
		}

		cfg, err := config.Get()
		if err != nil {
			return err
		}

		if cfg.Host == "" {
			fmt.Println("Not authenticated")
			fmt.Println("Use: gt auth login --host <url> --token <token>")
			return nil
		}

		fmt.Printf("Host: %s\n", cfg.Host)
		if cfg.Token != "" {
			fmt.Println("Token: set")
		} else {
			fmt.Println("Token: not set")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authStatusCmd)

	authLoginCmd.Flags().StringP("host", "H", "", "Gitea host URL")
	authLoginCmd.Flags().StringP("token", "t", "", "API token")
}
