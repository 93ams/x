package tests

import (
	"github.com/tilau2328/x/src/go/package/adaptor/data/cql/repo/ddl"
	"testing"
)

func TestNewTableRepo(t *testing.T) {
	tests := []struct {
		name string
		want *ddl.TableRepo
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
