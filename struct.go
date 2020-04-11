// Package groot is the root package containing the domain types and service contracts associated with the types.
package groot

type Struct struct {
	ID       int64  `orm:"pk;auto;column(id)" json:"-"`
	StructID string `orm:"column(struct_id)" json:"id"`
	Value    int    `orm:"column(value)" json:"value"`
}

func (s *Struct) TableName() string {
	return "struct"
}

type StructRepository interface {
	Create(StructRequest) (Struct, error)
	Struct(ID string) (Struct, error)
	Structs() ([]Struct, error)
	Update(Struct) error
	Delete(Struct) error
}

type StructRequest struct {
	Value int `json:"value"`
}
