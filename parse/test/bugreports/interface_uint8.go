// Code generated https://github.com/denkhaus/genny DO NOT EDIT.
// Any changes will be lost if this file is regenerated.
// see https://github.com/denkhaus/genny

package bugreports

type InterfaceUint8 interface {
	DoSomthingUint8()
}

// Call calls a method on an instance of generic interface.
// Targets github issue 49
func CallWithUint8(i InterfaceUint8) {
	i.DoSomthingUint8()
}
