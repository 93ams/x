package integration

import (
	"github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
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
