// Code generated https://github.com/denkhaus/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.
// see https://github.com/denkhaus/genny

package multipletypes

type StringIntMap map[string]int

func (m StringIntMap) Has(key string) bool {
	_, ok := m[key]
	return ok
}

func (m StringIntMap) Get(key string) int {
	return m[key]
}

func (m StringIntMap) Set(key string, value int) StringIntMap {
	m[key] = value
	return m
}
