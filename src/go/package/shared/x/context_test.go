package x

import (
	"context"
	"reflect"
	"testing"
)

type (
	TestStruct struct {
		val string
	}
	TestKey struct {
		val string
	}
)

func TestFromCtx(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		key     TestKey
		wantRet TestStruct
	}{
		{
			name:    "ineffable",
			ctx:     context.WithValue(context.Background(), TestKey{val: "asd"}, TestStruct{val: "dsa"}),
			key:     TestKey{val: "asd"},
			wantRet: TestStruct{val: "dsa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := FromCtx[TestKey, TestStruct](tt.ctx, tt.key); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("FromCtx() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
