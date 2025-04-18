package update

import (
	"context"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/spf13/afero"
	"github.com/skorpland/cli/internal/utils"
	"github.com/skorpland/cli/pkg/api"
)

func Run(ctx context.Context, branchId string, body api.UpdateBranchBody, fsys afero.Fs) error {
	resp, err := utils.GetPowerbase().V1UpdateABranchConfigWithResponse(ctx, branchId, body)
	if err != nil {
		return errors.Errorf("failed to update preview branch: %w", err)
	}
	if resp.JSON200 == nil {
		return errors.New("Unexpected error updating preview branch: " + string(resp.Body))
	}
	fmt.Println("Updated preview branch:", resp.JSON200.Id)
	return nil
}
