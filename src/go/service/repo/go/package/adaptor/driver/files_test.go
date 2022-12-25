package driver

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{
		{file: "./files.go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//CreateFile("./files_copy.go", lo.Must(ReadFile(tt.file)))
		})
	}
}
