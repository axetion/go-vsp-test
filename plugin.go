package main

//#include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	ifaceOK = iota
	ifaceFail
)

var description *C.char

//export Load
func Load(_ unsafe.Pointer, factory1 unsafe.Pointer, factory2 unsafe.Pointer) bool {
	fmt.Println("Go loading...")
	description = C.CString("GoTest")

	return true
}

//export Unload
func Unload(_ unsafe.Pointer) {
	fmt.Println("Go unloading...")
	C.free(unsafe.Pointer(description))
}

//export Description
func Description(_ unsafe.Pointer) *C.char {
	return description
}

//export PutInServer
func PutInServer(_ unsafe.Pointer, player unsafe.Pointer, name *C.char) {
	fmt.Printf("Go ClientPutInServer: %s\n", C.GoString(name))
}

//export CreateInterface
func CreateInterface(iface *C.char, ret *uint32) unsafe.Pointer {
	if C.GoString(iface) == "ISERVERPLUGINCALLBACKS003" {
		if ret != nil {
			*ret = ifaceOK
		}

		return unsafe.Pointer(vptr)
	}

	if ret != nil {
		*ret = ifaceFail
	}

	return nil
}

func main() {}
