// Code generated https://github.com/denkhaus/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.
// see https://github.com/denkhaus/genny

package bugreports

func ContainsString(slice []string, element string) bool {
	return false
}

// ContainsAllString targets github issue 36
func ContainsAllString(slice []string, other []string) bool {
	for _, e := range other {
		if !ContainsString(slice, e) {
			return false
		}
	}
	return true
}
