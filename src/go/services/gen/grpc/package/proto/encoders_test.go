package proto

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test__file_Encode(t *testing.T) {
	tests := []struct {
		name    string
		file    _file
		indent  int
		wantDst string
		wantErr error
	}{
		{
			name: "",
			file: File("test").
				Imports(
					Import("dsa", "asd").T,
				).
				Options(
					Option("dsa", "asd").T,
				).
				Extensions(
					Extension("test_ext",
						Field("asdField", TypeString, 0).T,
					).T,
				).
				Enums(
					Enum("testEnum",
						EnumElement("testEnumVal", 0).Comment("testEnum1").T,
						EnumElement("testEnumVal2", 1).Comment("testEnum2").T,
					).T,
				).
				Messages(
					Message("testMessage",
						Field("asdField", TypeString, 0).T,
						Field("asdField2", TypeUint64, 1).T,
					).T,
				).
				Services(
					Service("testService",
						Method("testMethod", "testInput", "testOutput").T,
					).T,
				).T,
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
