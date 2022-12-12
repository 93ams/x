//go:build tools

package _package

import (
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	_ "google.golang.org/protobuf/runtime/protoimpl"
)
