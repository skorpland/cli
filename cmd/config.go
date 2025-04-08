package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/skorpland/cli/internal/config/push"
	"github.com/skorpland/cli/internal/utils/flags"
)

var (
	configCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "config",
		Short:   "Manage Powerbase project configurations",
	}

	configPushCmd = &cobra.Command{
		Use:   "push",
		Short: "Pushes local config.toml to the linked project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return push.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}
)

func init() {
	configCmd.PersistentFlags().StringVar(&flags.ProjectRef, "project-ref", "", "Project ref of the Powerbase project.")
	configCmd.AddCommand(configPushCmd)
	rootCmd.AddCommand(configCmd)
}
