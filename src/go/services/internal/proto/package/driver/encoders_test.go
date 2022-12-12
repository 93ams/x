package driver_test

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"grpc/package/model"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		file    *FileBuilder
		indent  int
		wantDst string
		wantErr error
	}{
		{
			name: "sanity check",
			file: File("test").
				Imports(
					Import("dsa", "asd"),
				).
				Options(
					Option("dsa", "asd"),
				).
				Extensions(
					Extension("test_ext",
						Field("asdField", model.TypeString, 0),
					),
				).
				Enums(
					Enum("testEnum",
						EnumElement("testEnumVal", 0).Comment("testEnum1"),
						EnumElement("testEnumVal2", 1).Comment("testEnum2"),
					),
				).
				Messages(
					Message("testMessage",
						Field("asdField", model.TypeString, 0),
						Field("asdField2", model.TypeUint64, 1),
					),
				).
				Services(
					Service("testService",
						Method("testMethod", "testInput", "testOutput"),
					),
				),
			wantDst: `syntax = "proto3";

package test;

import "asd";

option dsa = "asd";

extend test_ext {
	string asdField = 0;
}

message testMessage {
	string asdField = 0;
	uint64 asdField2 = 1;
}

enum testEnum {
	testEnumVal = 0; // testEnum1
	testEnumVal2 = 1; // testEnum2
}

service testService {
	rpc testMethod(testInput) returns (testOutput);
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := &bytes.Buffer{}
			require.ErrorIs(t, tt.file.Encode(tt.indent, dst), tt.wantErr)
			require.Equal(t, tt.wantDst, dst.String())
		})
	}
}
