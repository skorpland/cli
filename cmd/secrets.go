package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/skorpland/cli/internal/secrets/list"
	"github.com/skorpland/cli/internal/secrets/set"
	"github.com/skorpland/cli/internal/secrets/unset"
	"github.com/skorpland/cli/internal/utils/flags"
)

var (
	secretsCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "secrets",
		Short:   "Manage Powerbase secrets",
	}

	secretsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all secrets on Powerbase",
		Long:  "List all secrets in the linked project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}

	secretsSetCmd = &cobra.Command{
		Use:   "set <NAME=VALUE> ...",
		Short: "Set a secret(s) on Powerbase",
		Long:  "Set a secret(s) to the linked Powerbase project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return set.Run(cmd.Context(), flags.ProjectRef, envFilePath, args, afero.NewOsFs())
		},
	}

	secretsUnsetCmd = &cobra.Command{
		Use:   "unset [NAME] ...",
		Short: "Unset a secret(s) on Powerbase",
		Long:  "Unset a secret(s) from the linked Powerbase project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return unset.Run(cmd.Context(), flags.ProjectRef, args, afero.NewOsFs())
		},
	}
)

func init() {
	secretsCmd.PersistentFlags().StringVar(&flags.ProjectRef, "project-ref", "", "Project ref of the Powerbase project.")
	secretsSetCmd.Flags().StringVar(&envFilePath, "env-file", "", "Read secrets from a .env file.")
	secretsCmd.AddCommand(secretsListCmd)
	secretsCmd.AddCommand(secretsSetCmd)
	secretsCmd.AddCommand(secretsUnsetCmd)
	rootCmd.AddCommand(secretsCmd)
}
