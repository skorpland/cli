package get

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-errors/errors"
	"github.com/skorpland/cli/internal/sso/internal/render"
	"github.com/skorpland/cli/internal/utils"
	"github.com/skorpland/cli/pkg/api"
)

func Run(ctx context.Context, ref, providerId, format string) error {
	resp, err := utils.GetPowerbase().V1GetASsoProviderWithResponse(ctx, ref, providerId)
	if err != nil {
		return err
	}

	if resp.JSON200 == nil {
		if resp.StatusCode() == http.StatusNotFound {
			return errors.Errorf("An identity provider with ID %q could not be found.", providerId)
		}

		return errors.New("Unexpected error fetching identity provider: " + string(resp.Body))
	}

	switch format {
	case utils.OutputMetadata:
		_, err := fmt.Println(*resp.JSON200.Saml.MetadataXml)
		return err
	case utils.OutputPretty:
		return render.SingleMarkdown(api.Provider{
			Id:        resp.JSON200.Id,
			Saml:      resp.JSON200.Saml,
			Domains:   resp.JSON200.Domains,
			CreatedAt: resp.JSON200.CreatedAt,
			UpdatedAt: resp.JSON200.UpdatedAt,
		})
	case utils.OutputEnv:
		return errors.Errorf("--output env flag is not supported")
	default:
		return utils.EncodeOutput(format, os.Stdout, resp.JSON200)
	}
}
