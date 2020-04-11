// Package inmemory is the in-memory implementation of StructRepository interface.
package inmemory

import (
	groot "beego-standard-layout"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type StructRepository map[string]groot.Struct

func NewStructRepository() StructRepository {
	// setup the initial structs in repository
	repo := map[string]groot.Struct{
		"id1": {StructID: "id1", Value: 1},
		"id2": {StructID: "id2", Value: 2},
	}

	return repo
}

func (s StructRepository) Create(req groot.StructRequest) (groot.Struct, error) {
	ID := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	newStruct := groot.Struct{StructID: ID, Value: req.Value}
	s[ID] = newStruct

	return newStruct, nil
}

func (s StructRepository) Struct(ID string) (groot.Struct, error) {
	if v, ok := s[ID]; ok {
		return v, nil
	}

	return groot.Struct{}, orm.ErrNoRows
}

func (s StructRepository) Structs() ([]groot.Struct, error) {
	if len(s) < 1 {
		return nil, orm.ErrNoRows
	}

	var structs []groot.Struct
	for _, v := range s {
		structs = append(structs, v)
	}

	return structs, nil
}

func (s StructRepository) Update(str groot.Struct) error {
	s[str.StructID] = str

	return nil
}

func (s StructRepository) Delete(str groot.Struct) error {
	delete(s, str.StructID)

	return nil
}
