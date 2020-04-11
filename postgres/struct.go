package postgres

import (
	groot "beego-standard-layout"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type StructRepository struct {
	DB orm.Ormer
}

func NewStructRepository(db orm.Ormer) StructRepository {
	return StructRepository{DB: db}
}

func (s StructRepository) Create(request groot.StructRequest) (groot.Struct, error) {
	strID := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	str := groot.Struct{StructID: strID, Value: request.Value}

	var err error
	str.ID, err = s.DB.Insert(&str)

	return str, err
}

func (s StructRepository) Struct(ID string) (groot.Struct, error) {
	str := groot.Struct{StructID: ID}
	err := s.DB.Read(&str, "StructID")

	return str, err
}

func (s StructRepository) Structs() ([]groot.Struct, error) {
	var structs []groot.Struct
	num, err := s.DB.QueryTable("struct").All(&structs)
	if num < 1 {
		return structs, orm.ErrNoRows
	}

	return structs, err
}

func (s StructRepository) Update(str groot.Struct) error {
	_, err := s.DB.Update(&str)

	return err
}

func (s StructRepository) Delete(str groot.Struct) error {
	_, err := s.DB.Delete(&str)

	return err
}
