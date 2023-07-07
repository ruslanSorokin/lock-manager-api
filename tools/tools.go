//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
	_ "golang.org/x/mod/gosumcheck"
	_ "golang.org/x/tools/cmd/goimports"
)
