package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/skorpland/cli/internal/services"
)

var (
	servicesCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "services",
		Short:   "Show versions of all Powerbase services",
		RunE: func(cmd *cobra.Command, args []string) error {
			return services.Run(cmd.Context(), afero.NewOsFs())
		},
	}
)

func init() {
	rootCmd.AddCommand(servicesCmd)
}
