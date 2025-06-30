package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/cryptexctl/teah/api"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	Long:  "Issue-related commands for repositories.",
}

var issueListCmd = &cobra.Command{
	Use:   "list <owner/repo>",
	Short: "List issues",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		state, _ := cmd.Flags().GetString("state")
		path := fmt.Sprintf("repos/%s/issues?state=%s", args[0], state)

		var issues []api.Issue
		if err := client.Get(path, &issues); err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NUMBER\tTITLE\tSTATE\tAUTHOR")
		for _, issue := range issues {
			fmt.Fprintf(w, "#%d\t%s\t%s\t%s\n", issue.Number, issue.Title, issue.State, issue.User.Login)
		}
		w.Flush()
		return nil
	},
}

var issueViewCmd = &cobra.Command{
	Use:   "view <owner/repo> <number>",
	Short: "View issue",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		number, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid issue number: %s", args[1])
		}

		path := fmt.Sprintf("repos/%s/issues/%d", args[0], number)

		var issue api.Issue
		if err := client.Get(path, &issue); err != nil {
			return err
		}

		fmt.Printf("Issue #%d: %s\n", issue.Number, issue.Title)
		fmt.Printf("State: %s\n", issue.State)
		fmt.Printf("Author: %s\n", issue.User.Login)
		fmt.Printf("URL: %s\n", issue.HTMLURL)
		fmt.Printf("\n%s\n", issue.Body)
		return nil
	},
}

var issueCreateCmd = &cobra.Command{
	Use:   "create <owner/repo>",
	Short: "Create issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")

		if title == "" {
			return fmt.Errorf("title is required, use --title")
		}

		req := api.CreateIssueRequest{
			Title: title,
			Body:  body,
		}

		path := fmt.Sprintf("repos/%s/issues", args[0])

		var issue api.Issue
		if err := client.Post(path, req, &issue); err != nil {
			return err
		}

		fmt.Printf("Issue created #%d: %s\n", issue.Number, issue.HTMLURL)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
	issueCmd.AddCommand(issueListCmd)
	issueCmd.AddCommand(issueViewCmd)
	issueCmd.AddCommand(issueCreateCmd)

	issueListCmd.Flags().StringP("state", "s", "open", "State: open, closed, all")
	issueCreateCmd.Flags().StringP("title", "t", "", "Issue title")
	issueCreateCmd.Flags().StringP("body", "b", "", "Issue description")
}
