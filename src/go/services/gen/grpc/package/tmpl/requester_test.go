package tmpl

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestRequester(t *testing.T) {
	tests := []struct {
		name    string
		props   RequesterProps
		wantW   string
		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			require.ErrorIs(t, tt.wantErr, Requester(w, tt.props))
			require.Equal(t, tt.wantW, w.String())
			log.Println(w.String())
		})
	}
}
