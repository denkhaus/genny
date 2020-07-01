package unexported

import (
	"github.com/denkhaus/genny/generic"
)

type tag generic.Type
type TypeName generic.Type

type TypeNameType struct {
	Field int `json:"tag"`
}
