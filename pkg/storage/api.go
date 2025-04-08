package storage

import "github.com/skorpland/cli/pkg/fetcher"

type StorageAPI struct {
	*fetcher.Fetcher
}

const PAGE_LIMIT = 100
