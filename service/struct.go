// Package service is for service/operation on domain types.
package service

import (
	groot "beego-standard-layout"
	"beego-standard-layout/response"
)

type StructService struct {
	repo groot.StructRepository
}

func NewStructService(repo groot.StructRepository) StructService {
	return StructService{repo: repo}
}

func (s StructService) CreateStruct(request groot.StructRequest) response.Response {
	str, err := s.repo.Create(request)
	if err != nil {
		return response.Error(err.Error())
	}

	return response.Success(str)
}
