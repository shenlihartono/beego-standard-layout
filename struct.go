// Package groot is the root package containing the domain types and service contracts associated with the types.
package groot

type Struct struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type StructRepository interface {
	Create(StructRequest) (Struct, error)
	Struct(ID string) (Struct, error)
	Structs() ([]Struct, error)
	Update(Struct) error
	Delete(ID string) error
}

type StructRequest struct {
	Value int `json:"value"`
}
