package tmpl

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name    string
		props   HandlerProps
		wantW   string
		wantErr error
	}{
		{
			props: HandlerProps{
				Pkg:      "test_pkg",
				Name:     "test",
				Provider: "provider",
				Imports: map[string]string{
					"log": "",
				},
				Methods: []MethodProps{{
					Name:    "test_method",
					Return:  "something_else",
					Request: "something",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			require.ErrorIs(t, tt.wantErr, Handler(w, tt.props))
			require.Equal(t, tt.wantW, w.String())
			log.Println(w.String())
		})
	}
}
