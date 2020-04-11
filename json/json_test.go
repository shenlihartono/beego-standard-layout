package json

import (
	groot "beego-standard-layout"
	"fmt"
	"github.com/go-test/deep"
	"testing"
)

var emptyStruct groot.Struct

func TestConvertRequest(t *testing.T) {
	tests := []struct {
		name       string
		reqByte    []byte
		reqStruct  groot.Struct
		wantErr    bool
		wantStruct groot.Struct
	}{
		{
			name:       "attempt to convert empty request body",
			reqByte:    []byte(`{}`),
			reqStruct:  emptyStruct,
			wantErr:    true,
			wantStruct: emptyStruct,
		},
		{
			name:       "attempt to convert nil request body",
			reqStruct:  emptyStruct,
			wantErr:    true,
			wantStruct: emptyStruct,
		},
		{
			name:       "attempt to convert invalid data type",
			reqByte:    []byte(`{"id": "123", "value": xyz}`),
			reqStruct:  emptyStruct,
			wantErr:    true,
			wantStruct: emptyStruct,
		},
		{
			name:       "attempt to convert invalid json ",
			reqByte:    []byte(`sorry not json`),
			reqStruct:  emptyStruct,
			wantErr:    true,
			wantStruct: emptyStruct,
		},
		{
			name:       "attempt to convert json with length < 10 characters",
			reqByte:    []byte(`{"a": 0}`),
			reqStruct:  emptyStruct,
			wantErr:    true,
			wantStruct: emptyStruct,
		},
		{
			name:       "success story convert valid json",
			reqByte:    []byte(`{"id": "xyz", "value": 123}`),
			reqStruct:  emptyStruct,
			wantErr:    false,
			wantStruct: groot.Struct{StructID: "xyz", Value: 123},
		},
	}

	for i, tc := range tests {
		no := i + 1
		t.Run(fmt.Sprintf("Test no %d %s", no, tc.name), func(t *testing.T) {
			err := ConvertRequest(tc.reqByte, &tc.reqStruct)
			var haveErr, isEmpty bool
			if err != nil {
				haveErr = true
				isEmpty = true
			}

			if haveErr != tc.wantErr {
				t.Errorf("have error: %v, want error: %v", haveErr, tc.wantErr)
			}

			if !isEmpty {
				if diff := deep.Equal(tc.reqStruct, tc.wantStruct); diff != nil {
					t.Error(diff)
				}
			}

		})
	}
}
