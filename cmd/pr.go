package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/cryptexctl/teah/api"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Manage pull requests",
	Long:  "Pull request related commands.",
}

var prListCmd = &cobra.Command{
	Use:   "list <owner/repo>",
	Short: "List pull requests",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		state, _ := cmd.Flags().GetString("state")
		path := fmt.Sprintf("repos/%s/pulls?state=%s", args[0], state)

		var prs []api.PullRequest
		if err := client.Get(path, &prs); err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NUMBER\tTITLE\tSTATE\tAUTHOR\tHEAD->BASE")
		for _, pr := range prs {
			fmt.Fprintf(w, "#%d\t%s\t%s\t%s\t%s->%s\n",
				pr.Number, pr.Title, pr.State, pr.User.Login, pr.Head.Ref, pr.Base.Ref)
		}
		w.Flush()
		return nil
	},
}

var prViewCmd = &cobra.Command{
	Use:   "view <owner/repo> <number>",
	Short: "View pull request",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		number, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid PR number: %s", args[1])
		}

		path := fmt.Sprintf("repos/%s/pulls/%d", args[0], number)

		var pr api.PullRequest
		if err := client.Get(path, &pr); err != nil {
			return err
		}

		fmt.Printf("Pull Request #%d: %s\n", pr.Number, pr.Title)
		fmt.Printf("State: %s\n", pr.State)
		fmt.Printf("Author: %s\n", pr.User.Login)
		fmt.Printf("Branches: %s -> %s\n", pr.Head.Ref, pr.Base.Ref)
		fmt.Printf("URL: %s\n", pr.HTMLURL)
		fmt.Printf("\n%s\n", pr.Body)
		return nil
	},
}

var prCreateCmd = &cobra.Command{
	Use:   "create <owner/repo>",
	Short: "Create pull request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		head, _ := cmd.Flags().GetString("head")
		base, _ := cmd.Flags().GetString("base")

		if title == "" {
			return fmt.Errorf("title is required, use --title")
		}
		if head == "" {
			return fmt.Errorf("head branch is required, use --head")
		}
		if base == "" {
			base = "main"
		}

		req := api.CreatePRRequest{
			Title: title,
			Body:  body,
			Head:  head,
			Base:  base,
		}

		path := fmt.Sprintf("repos/%s/pulls", args[0])

		var pr api.PullRequest
		if err := client.Post(path, req, &pr); err != nil {
			return err
		}

		fmt.Printf("Pull request created #%d: %s\n", pr.Number, pr.HTMLURL)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.AddCommand(prListCmd)
	prCmd.AddCommand(prViewCmd)
	prCmd.AddCommand(prCreateCmd)

	prListCmd.Flags().StringP("state", "s", "open", "State: open, closed, all")
	prCreateCmd.Flags().StringP("title", "t", "", "PR title")
	prCreateCmd.Flags().StringP("body", "b", "", "PR description")
	prCreateCmd.Flags().String("head", "", "Head branch")
	prCreateCmd.Flags().String("base", "main", "Base branch")
}
