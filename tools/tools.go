//go:build tools

package tools

import (
	// Import all required tools so that they can be version controlled via go.mod & go.sum.
	// These imports are tracked a sub package so as to not impact the dependencies of clients of
	// the library.
	// This should be migrated directly to go.mod when the following is complete:
	// https://github.com/golang/go/issues/48429
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "gotest.tools/gotestsum"
)
