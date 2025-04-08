package cmd

import (
	"os"
	"os/signal"

	env "github.com/Netflix/go-env"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/skorpland/cli/internal/status"
	"github.com/skorpland/cli/internal/utils"
)

var (
	override []string
	names    status.CustomName

	statusCmd = &cobra.Command{
		GroupID: groupLocalDev,
		Use:     "status",
		Short:   "Show status of local Powerbase containers",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			es, err := env.EnvironToEnvSet(override)
			if err != nil {
				return err
			}
			return env.Unmarshal(es, &names)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, _ := signal.NotifyContext(cmd.Context(), os.Interrupt)
			return status.Run(ctx, names, utils.OutputFormat.Value, afero.NewOsFs())
		},
		Example: `  powerbase status -o env --override-name api.url=NEXT_PUBLIC_POWERBASE_URL
  powerbase status -o json`,
	}
)

func init() {
	flags := statusCmd.Flags()
	flags.StringSliceVar(&override, "override-name", []string{}, "Override specific variable names.")
	rootCmd.AddCommand(statusCmd)
}
