package service

import (
	groot "beego-standard-layout"
	"beego-standard-layout/mock"
	"errors"
	"fmt"
	"github.com/go-test/deep"
	"testing"
)

var (
	createSuccessResult     = groot.Struct{ID: "abc123", Value: 500}
	createSuccessResponse   = Response{Status: 200, Result: createSuccessResult}
	createSuccessRepository = mock.StructRepository{TheStruct: createSuccessResult}

	createFailedResponse   = Response{Status: 500, Result: "internal server error"}
	createFailedRepository = mock.StructRepository{ErrCreate: errors.New("failed create struct")}

	createRequest = groot.StructRequest{}
)

func TestCreateStruct(t *testing.T) {
	tests := []struct {
		name     string
		repo     groot.StructRepository
		request  groot.StructRequest
		wantResp Response
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

var (
	findSuccessResult         = groot.Struct{ID: "ABC123", Value: 500}
	findSuccessRepository     = mock.StructRepository{TheStruct: findSuccessResult}
	findStructSuccessResponse = Response{Status: 200, Result: findSuccessResult}

	failedFindNotFoundRepository = mock.StructRepository{ErrStruct: errStructNotFound}
	findStructNotFoundResponse   = Response{Status: 404, Result: "struct not found"}

	failedFindOtherErrRepository = mock.StructRepository{ErrStruct: errors.New("error find struct")}
	findStructOtherErrResponse   = Response{Status: 500, Result: "internal server error"}
)

func TestFindStruct(t *testing.T) {
	tests := []struct {
		name     string
		repo     groot.StructRepository
		request  string
		wantResp Response
	}{
		{
			name:     "success find 1 struct",
			repo:     findSuccessRepository,
			request:  "ABC123",
			wantResp: findStructSuccessResponse,
		},
		{
			name:     "failed find 1 struct, not found",
			repo:     failedFindNotFoundRepository,
			request:  "ABC123",
			wantResp: findStructNotFoundResponse,
		},
		{
			name:     "failed find 1 struct, other error",
			repo:     failedFindOtherErrRepository,
			request:  "ABC123",
			wantResp: findStructOtherErrResponse,
		},
	}

	for i, tc := range tests {
		no := i + 1
		t.Run(fmt.Sprintf("Test no %d %s", no, tc.name), func(t *testing.T) {
			svc := NewStructService(tc.repo)
			resp := svc.Struct(tc.request)

			if diff := deep.Equal(resp, tc.wantResp); diff != nil {
				t.Error(diff)
			}
		})
	}
}

var (
	findSuccessResults = []groot.Struct{
		{ID: "one", Value: 1},
		{ID: "two", Value: 2},
	}
	findStructsSuccessRepository = mock.StructRepository{TheStructs: findSuccessResults}
	findStructsSuccessResponse   = Response{Status: 200, Result: findSuccessResults}
)

func TestFindStructs(t *testing.T) {
	tests := []struct {
		name     string
		repo     groot.StructRepository
		wantResp Response
	}{
		{
			name:     "success find all struct",
			repo:     findStructsSuccessRepository,
			wantResp: findStructsSuccessResponse,
		},
		{
			name:     "failed find all struct, not found",
			repo:     failedFindNotFoundRepository,
			wantResp: findStructNotFoundResponse,
		},
		{
			name:     "failed find all struct, other error",
			repo:     failedFindOtherErrRepository,
			wantResp: findStructOtherErrResponse,
		},
	}

	for i, tc := range tests {
		no := i + 1
		t.Run(fmt.Sprintf("Test no %d %s", no, tc.name), func(t *testing.T) {
			svc := NewStructService(tc.repo)
			resp := svc.Structs()

			if diff := deep.Equal(resp, tc.wantResp); diff != nil {
				t.Error(diff)
			}
		})
	}
}

var (
	updateSuccessRepository = mock.StructRepository{
		TheStruct: groot.Struct{ID: "ABC123", Value: 1500},
	}
	updateSuccessResult   = groot.Struct{ID: "ABC123", Value: 2000}
	updateSuccessResponse = Response{Status: StatusOK, Result: updateSuccessResult}

	updateFailedOtherErrorRepository = mock.StructRepository{ErrUpdate: errors.New("error update struct")}
	updateFailedResponse             = Response{Status: 500, Result: "internal server error"}

	updateRequest = groot.StructRequest{Value: 2000}
)

func TestUpdateStruct(t *testing.T) {
	tests := []struct {
		name     string
		repo     groot.StructRepository
		structId string
		request  groot.StructRequest
		wantResp Response
	}{
		{
			name:     "failed find struct, not found",
			repo:     failedFindNotFoundRepository,
			structId: "ABC123",
			request:  updateRequest,
			wantResp: findStructNotFoundResponse,
		},
		{
			name:     "failed find struct, other error",
			repo:     failedFindOtherErrRepository,
			structId: "ABC123",
			request:  updateRequest,
			wantResp: findStructOtherErrResponse,
		},
		{
			name:     "failed update struct, internal error",
			repo:     updateFailedOtherErrorRepository,
			structId: "ABC123",
			request:  updateRequest,
			wantResp: updateFailedResponse,
		},
		{
			name:     "success update struct",
			repo:     updateSuccessRepository,
			structId: "ABC123",
			request:  updateRequest,
			wantResp: updateSuccessResponse,
		},
	}

	for i, tc := range tests {
		no := i + 1
		t.Run(fmt.Sprintf("Test no %d %s", no, tc.name), func(t *testing.T) {
			svc := NewStructService(tc.repo)
			resp := svc.UpdateStruct(tc.structId, tc.request)

			if diff := deep.Equal(resp, tc.wantResp); diff != nil {
				t.Error(diff)
			}
		})
	}
}
