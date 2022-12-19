package tests

import (
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
