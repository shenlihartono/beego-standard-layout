package response

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
			want:   Response{"status": 200, "result": "hello"},
		},
		{
			name:   "generate response failed",
			status: 400,
			result: "world",
			want:   Response{"status": 400, "result": "world"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			have := New(tc.status, tc.result)

			if diff := deep.Equal(have, tc.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
