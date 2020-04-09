package service

import (
	"github.com/go-test/deep"
	"testing"
)

func TestResponse_New(t *testing.T) {
	tests := []struct {
		name   string
		status int
		result interface{}
		want   Response
	}{
		{
			name:   "generate response success",
			status: 200,
			result: "hello",
			want:   Response{Status: 200, Result: "hello"},
		},
		{
			name:   "generate response failed",
			status: 400,
			result: "world",
			want:   Response{Status: 400, Result: "world"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			have := NewResponse(tc.status, tc.result)

			if diff := deep.Equal(have, tc.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
