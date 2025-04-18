package function

import (
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/skorpland/cli/pkg/api"
	"github.com/skorpland/cli/pkg/config"
)

type MockBundler struct {
}

func (b *MockBundler) Bundle(ctx context.Context, slug, entrypoint, importMap string, staticFiles []string, output io.Writer) (api.FunctionDeployMetadata, error) {
	if staticFiles == nil {
		staticFiles = []string{}
	}
	return api.FunctionDeployMetadata{
		Name:           &slug,
		EntrypointPath: entrypoint,
		ImportMapPath:  &importMap,
		StaticPatterns: &staticFiles,
	}, nil
}

const (
	mockApiHost = "https://api.powerbase.club"
	mockProject = "test-project"
)

func TestUpsertFunctions(t *testing.T) {
	apiClient, err := api.NewClientWithResponses(mockApiHost)
	require.NoError(t, err)
	client := NewEdgeRuntimeAPI(mockProject, *apiClient, func(era *EdgeRuntimeAPI) {
		era.eszip = &MockBundler{}
	})

	t.Run("throws error on network failure", func(t *testing.T) {
		// Setup mock api
		defer gock.OffAll()
		gock.New(mockApiHost).
			Get("/v1/projects/" + mockProject + "/functions").
			ReplyError(errors.New("network error"))
		// Run test
		err := client.UpsertFunctions(context.Background(), nil)
		// Check error
		assert.ErrorContains(t, err, "network error")
	})

	t.Run("throws error on service unavailable", func(t *testing.T) {
		// Setup mock api
		defer gock.OffAll()
		gock.New(mockApiHost).
			Get("/v1/projects/" + mockProject + "/functions").
			Reply(http.StatusServiceUnavailable)
		// Run test
		err := client.UpsertFunctions(context.Background(), nil)
		// Check error
		assert.ErrorContains(t, err, "unexpected list functions status 503:")
	})

	t.Run("retries on create failure", func(t *testing.T) {
		// Setup mock api
		defer gock.OffAll()
		gock.New(mockApiHost).
			Get("/v1/projects/" + mockProject + "/functions").
			Reply(http.StatusOK).
			JSON([]api.FunctionResponse{})
		gock.New(mockApiHost).
			Post("/v1/projects/" + mockProject + "/functions").
			ReplyError(errors.New("network error"))
		gock.New(mockApiHost).
			Post("/v1/projects/" + mockProject + "/functions").
			Reply(http.StatusServiceUnavailable)
		gock.New(mockApiHost).
			Post("/v1/projects/" + mockProject + "/functions").
			Reply(http.StatusCreated).
			JSON(api.FunctionResponse{Slug: "test"})
		// Run test
		err := client.UpsertFunctions(context.Background(), config.FunctionConfig{
			"test": {},
		})
		// Check error
		assert.NoError(t, err)
	})

	t.Run("retries on update failure", func(t *testing.T) {
		// Setup mock api
		defer gock.OffAll()
		gock.New(mockApiHost).
			Get("/v1/projects/" + mockProject + "/functions").
			Reply(http.StatusOK).
			JSON([]api.FunctionResponse{{Slug: "test"}})
		gock.New(mockApiHost).
			Patch("/v1/projects/" + mockProject + "/functions/test").
			ReplyError(errors.New("network error"))
		gock.New(mockApiHost).
			Patch("/v1/projects/" + mockProject + "/functions/test").
			Reply(http.StatusServiceUnavailable)
		gock.New(mockApiHost).
			Patch("/v1/projects/" + mockProject + "/functions/test").
			Reply(http.StatusOK).
			JSON(api.FunctionResponse{Slug: "test"})
		// Run test
		err := client.UpsertFunctions(context.Background(), config.FunctionConfig{
			"test": {},
		})
		// Check error
		assert.NoError(t, err)
	})
}
