// Code generated https://github.com/denkhaus/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.
// see https://github.com/denkhaus/genny

package bugreports

// ObjInt is the struct used for tests.
type ObjInt struct {
	v int
}

// NewInts calls new on ObjInt and instantiates slice.
// Targets github issues #36 and #49
func NewInts() (*ObjInt, []ObjInt) {
	n := new(ObjInt)
	m := make([]ObjInt, 0)
	return n, m
}
