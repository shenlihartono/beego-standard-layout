package response

import (
	"github.com/go-test/deep"
	"testing"
)

func TestResponseSuccess(t *testing.T) {
	tests := []struct {
		name string
		body interface{}
		want Response
	}{
		{
			name: "generate response success",
			body: "hello",
			want: Response{"result": "hello"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			have := Success(tc.body)

			if diff := deep.Equal(have, tc.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestResponseError(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    Response
	}{
		{
			name:    "generate response error",
			message: "something wrong",
			want:    Response{"error": "something wrong"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			have := Error(tc.message)

			if diff := deep.Equal(have, tc.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
