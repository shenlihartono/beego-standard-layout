package service

import (
	groot "beego-standard-layout"
	"beego-standard-layout/mock"
	"beego-standard-layout/response"
	"errors"
	"fmt"
	"github.com/go-test/deep"
	"testing"
)

var (
	createSuccessResult     = groot.Struct{ID: "abc123", Value: 500}
	createSuccessResponse   = response.Response{"result": createSuccessResult}
	createSuccessRepository = mock.StructRepository{TheStruct: createSuccessResult}

	createFailedResponse   = response.Response{"error": "failed create struct"}
	createFailedRepository = mock.StructRepository{ErrCreate: errors.New("failed create struct")}

	createRequest = groot.StructRequest{}
)

func TestCreateStruct(t *testing.T) {
	tests := []struct {
		name     string
		repo     groot.StructRepository
		request  groot.StructRequest
		wantResp response.Response
	}{
		{
			name:     "success create struct",
			repo:     createSuccessRepository,
			request:  createRequest,
			wantResp: createSuccessResponse,
		},
		{
			name:     "failed create struct",
			repo:     createFailedRepository,
			request:  createRequest,
			wantResp: createFailedResponse,
		},
	}

	for i, tc := range tests {
		no := i + 1
		t.Run(fmt.Sprintf("Test no %d %s", no, tc.name), func(t *testing.T) {
			svc := NewStructService(tc.repo)
			resp := svc.CreateStruct(tc.request)

			if diff := deep.Equal(resp, tc.wantResp); diff != nil {
				t.Error(diff)
			}
		})
	}
}
