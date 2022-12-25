package x

import (
	"errors"
	"log"
	"testing"
)

func TestParallelTry(t *testing.T) {
	type args[T any] struct {
		fn   func(T) error
		args []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantErr bool
	}
	tests := []testCase[string]{
		{
			args: args[string]{
				fn: func(in string) error {
					log.Println(in)
					return errors.New("err")
				},
				args: []string{"foo", "bar"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(ParallelTry(tt.args.fn, tt.args.args))
		})
	}
}
