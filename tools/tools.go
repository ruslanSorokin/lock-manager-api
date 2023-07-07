//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
