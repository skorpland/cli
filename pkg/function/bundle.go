package function

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/go-errors/errors"
	"github.com/skorpland/cli/pkg/api"
	"github.com/skorpland/cli/pkg/cast"
)

type nativeBundler struct {
	tempDir string
	fsys    fs.FS
}

func NewNativeBundler(tempDir string, fsys fs.FS) EszipBundler {
	return &nativeBundler{
		tempDir: tempDir,
		fsys:    fsys,
	}
}

// Use a package private variable to allow testing without gosec complaining about G204
var edgeRuntimeBin = "edge-runtime"

func (b *nativeBundler) Bundle(ctx context.Context, slug, entrypoint, importMap string, staticFiles []string, output io.Writer) (api.FunctionDeployMetadata, error) {
	meta := NewMetadata(slug, entrypoint, importMap, staticFiles)
	outputPath := filepath.Join(b.tempDir, slug+".eszip")
	// TODO: make edge runtime write to stdout
	args := []string{"bundle", "--entrypoint", entrypoint, "--output", outputPath}
	if len(importMap) > 0 {
		args = append(args, "--import-map", importMap)
	}
	for _, staticFile := range staticFiles {
		args = append(args, "--static", staticFile)
	}
	cmd := exec.CommandContext(ctx, edgeRuntimeBin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return meta, errors.Errorf("failed to bundle function: %w", err)
	}
	defer os.Remove(outputPath)
	// Compress the output
	eszipBytes, err := b.fsys.Open(outputPath)
	if err != nil {
		return meta, errors.Errorf("failed to open eszip: %w", err)
	}
	defer eszipBytes.Close()
	return meta, Compress(eszipBytes, output)
}

const compressedEszipMagicID = "EZBR"

func Compress(r io.Reader, w io.Writer) error {
	if _, err := fmt.Fprint(w, compressedEszipMagicID); err != nil {
		return errors.Errorf("failed to append magic id: %w", err)
	}
	brw := brotli.NewWriter(w)
	defer brw.Close()
	if _, err := io.Copy(brw, r); err != nil {
		return errors.Errorf("failed to compress eszip: %w", err)
	}
	return nil
}

func NewMetadata(slug, entrypoint, importMap string, staticFiles []string) api.FunctionDeployMetadata {
	meta := api.FunctionDeployMetadata{
		Name:           &slug,
		EntrypointPath: toFileURL(entrypoint),
	}
	if len(importMap) > 0 {
		meta.ImportMapPath = cast.Ptr(toFileURL(importMap))
	}
	files := make([]string, len(staticFiles))
	for i, sf := range staticFiles {
		files[i] = toFileURL(sf)
	}
	meta.StaticPatterns = &files
	return meta
}

func toFileURL(hostPath string) string {
	absHostPath, err := filepath.Abs(hostPath)
	if err != nil {
		return hostPath
	}
	// Convert to unix path because edge runtime only supports linux
	unixPath := toUnixPath(absHostPath)
	parsed := url.URL{Scheme: "file", Path: unixPath}
	return parsed.String()
}

func toUnixPath(absHostPath string) string {
	prefix := filepath.VolumeName(absHostPath)
	unixPath := filepath.ToSlash(absHostPath)
	return strings.TrimPrefix(unixPath, prefix)
}
