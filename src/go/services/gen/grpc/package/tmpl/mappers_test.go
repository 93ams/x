package tmpl

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMappers(t *testing.T) {
	tests := []struct {
		name    string
		props   MappersProps
		wantW   string
		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			require.ErrorIs(t, tt.wantErr, Mappers(w, tt.props))
			require.Equal(t, tt.wantW, w.String())
		})
	}
}
