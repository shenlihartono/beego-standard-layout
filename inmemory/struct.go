// Package inmemory is the in-memory implementation of StructRepository interface.
package inmemory

import (
	groot "beego-standard-layout"
	"errors"
	"strconv"
	"strings"
	"time"
)

var errStructNotFound = errors.New("struct not found")

type StructRepository map[string]groot.Struct

func NewStructRepository() StructRepository {
	// setup the initial structs in repository
	repo := map[string]groot.Struct{
		"id1": {"id1", 1},
		"id2": {"id2", 2},
	}

	return repo
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

	return groot.Struct{}, errStructNotFound
}

func (s StructRepository) Structs() ([]groot.Struct, error) {
	if len(s) < 1 {
		return nil, errStructNotFound
	}

	var structs []groot.Struct
	for _, v := range s {
		structs = append(structs, v)
	}

	return structs, nil
}

func (s StructRepository) Update(req groot.Struct) error {
	s[req.ID] = req

	return nil
}

func (s StructRepository) Delete(ID string) {
	delete(s, ID)
}
