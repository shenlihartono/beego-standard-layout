// Package inmemory is the in-memory implementation of StructRepository interface.
package inmemory

import (
	groot "beego-standard-layout"
	"errors"
	"strconv"
	"strings"
	"time"
)

type StructRepository map[string]groot.Struct

func NewStructRepository() StructRepository {
	s := map[string]groot.Struct{
		"id1": {"id1", 1},
		"id2": {"id2", 2},
	}

	return s
}

func (s StructRepository) Create(req groot.StructRequest) (groot.Struct, error) {
	ID := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	newStruct := groot.Struct{ID: ID, Value: req.Value}
	s[ID] = newStruct

	return newStruct, nil
}

func (s StructRepository) Struct(ID string) (groot.Struct, error) {
	if v, ok := s[ID]; ok {
		return v, nil
	}

	return groot.Struct{}, errors.New("struct not found")
}

func (s StructRepository) Structs() ([]groot.Struct, error) {
	if len(s) < 1 {
		return nil, errors.New("no structs available")
	}

	var structs []groot.Struct
	for _, v := range s {
		structs = append(structs, v)
	}

	return structs, nil
}

func (s StructRepository) Update(ID string, req groot.StructRequest) error {
	// do find by ID and update/replace the value
	if _, ok := s[ID]; ok {
		s[ID] = groot.Struct{ID: ID, Value: req.Value}
		return nil
	}

	return errors.New("struct not found")
}

func (s StructRepository) Delete(ID string) {
	delete(s, ID)
}
