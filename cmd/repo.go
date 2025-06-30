package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cryptexctl/gt/api"
	"github.com/cryptexctl/gt/internal/config"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage repositories",
	Long:  "Repository-related commands for Gitea.",
}

var repoListCmd = &cobra.Command{
	Use:   "list [user]",
	Short: "List repositories",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		var path string
		if len(args) == 0 {
			path = "user/repos"
		} else {
			path = fmt.Sprintf("users/%s/repos", args[0])
		}

		var repos []api.Repository
		if err := client.Get(path, &repos); err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tDESCRIPTION\tPRIVATE\tFORK")
		for _, repo := range repos {
			private := ""
			if repo.Private {
				private = "yes"
			}
			fork := ""
			if repo.Fork {
				fork = "yes"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", repo.FullName, repo.Description, private, fork)
		}
		w.Flush()
		return nil
	},
}

var repoCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := getAPIClient()
		if err != nil {
			return err
		}

		description, _ := cmd.Flags().GetString("description")
		private, _ := cmd.Flags().GetBool("private")
		autoInit, _ := cmd.Flags().GetBool("auto-init")

		req := api.CreateRepoRequest{
			Name:        args[0],
			Description: description,
			Private:     private,
			AutoInit:    autoInit,
		}

		var repo api.Repository
		if err := client.Post("user/repos", req, &repo); err != nil {
			return err
		}

		fmt.Printf("Repository created: %s\n", repo.HTMLURL)
		return nil
	},
}

func getAPIClient() (*api.Client, error) {
	config.Init()

	host := rootCmd.Flag("host").Value.String()
	token := rootCmd.Flag("token").Value.String()

	if host == "" {
		host = config.GetHost()
	}
	if token == "" {
		token = config.GetToken()
	}

	if host == "" {
		return nil, fmt.Errorf("Gitea host not specified. Use --host or set in config")
	}
	if token == "" {
		return nil, fmt.Errorf("Gitea token not specified. Use --token or set in config")
	}

	return api.NewClient(host, token), nil
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoCreateCmd)

	repoCreateCmd.Flags().StringP("description", "d", "", "Repository description")
	repoCreateCmd.Flags().BoolP("private", "p", false, "Private repository")
	repoCreateCmd.Flags().Bool("auto-init", true, "Auto-create README")
}
