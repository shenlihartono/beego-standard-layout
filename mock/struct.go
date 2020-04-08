// Package mock contains the mock implementation of Struct Repository.
package mock

import groot "beego-standard-layout"

type StructRepository struct {
	ErrCreate error
	ErrStruct error
	ErrUpdate error

	TheStruct  groot.Struct
	TheStructs []groot.Struct
}

func (s StructRepository) Create(groot.StructRequest) (groot.Struct, error) {
	return s.TheStruct, s.ErrCreate
}

func (s StructRepository) Struct(ID string) (groot.Struct, error) {
	return s.TheStruct, s.ErrStruct
}

func (s StructRepository) Structs() ([]groot.Struct, error) {
	return s.TheStructs, s.ErrStruct
}

func (s StructRepository) Update(ID string, req groot.StructRequest) error {
	return s.ErrUpdate
}

func (s StructRepository) Delete(ID string) {}
