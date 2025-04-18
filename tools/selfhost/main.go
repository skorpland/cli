package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/google/go-github/v62/github"
	"github.com/skorpland/cli/internal/utils"
	"github.com/skorpland/cli/pkg/config"
	"github.com/skorpland/cli/tools/shared"
	"gopkg.in/yaml.v3"
)

const (
	POWERBASE_REPO  = "powerbase"
	POWERBASE_OWNER = "powerbase"
)

func main() {
	branch := "cli/latest"
	if len(os.Args) > 1 {
		branch = os.Args[1]
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	if err := updateSelfHosted(ctx, branch); err != nil {
		log.Fatalln(err)
	}
}

type ComposeService struct {
	Image string `yaml:"image,omitempty"`
}

type ComposeFile struct {
	Services map[string]ComposeService `yaml:"services,omitempty"`
}

func updateSelfHosted(ctx context.Context, branch string) error {
	client := utils.GetGitHubClient(ctx)
	master := "master"
	if err := shared.CreateGitBranch(ctx, client, POWERBASE_OWNER, POWERBASE_REPO, branch, master); err != nil {
		return err
	}
	stable := getStableVersions()
	if err := updateComposeVersion(ctx, client, "docker/docker-compose.yml", branch, stable); err != nil {
		return err
	}
	pr := github.NewPullRequest{
		Title: github.String("chore: update self-hosted image versions"),
		Head:  &branch,
		Base:  &master,
	}
	return shared.CreatePullRequest(ctx, client, POWERBASE_OWNER, POWERBASE_REPO, pr)
}

func getStableVersions() map[string]string {
	images := append(config.Images.Services(), config.Images.Pg15)
	result := make(map[string]string, len(images))
	for _, img := range images {
		parts := strings.Split(img, ":")
		key := strings.TrimPrefix(parts[0], "library/")
		result[key] = parts[1]
	}
	return result
}

func updateComposeVersion(ctx context.Context, client *github.Client, path, ref string, stable map[string]string) error {
	fmt.Fprintln(os.Stderr, "Parsing file:", path)
	opts := github.RepositoryContentGetOptions{Ref: "heads/" + ref}
	file, _, _, err := client.Repositories.GetContents(ctx, POWERBASE_OWNER, POWERBASE_REPO, path, &opts)
	if err != nil {
		return err
	}
	content, err := base64.StdEncoding.DecodeString(*file.Content)
	if err != nil {
		return err
	}
	var data ComposeFile
	if err := yaml.Unmarshal(content, &data); err != nil {
		return err
	}
	updated := false
	for _, v := range data.Services {
		parts := strings.Split(v.Image, ":")
		if version, ok := stable[parts[0]]; ok && parts[1] != version {
			fmt.Fprintf(os.Stderr, "Updating %s: %s => %s\n", parts[0], parts[1], version)
			image := strings.Join([]string{parts[0], version}, ":")
			content = bytes.ReplaceAll(content, []byte(v.Image), []byte(image))
			updated = true
		}
	}
	if !updated {
		fmt.Fprintln(os.Stderr, "All images are up to date.")
		return nil
	}
	message := "chore: update image versions for " + path
	commit := github.RepositoryContentFileOptions{
		Message: &message,
		Content: content,
		SHA:     file.SHA,
		Branch:  &ref,
	}
	resp, _, err := client.Repositories.UpdateFile(ctx, POWERBASE_OWNER, POWERBASE_REPO, path, &commit)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Committed changes to", *resp.SHA)
	return nil
}
