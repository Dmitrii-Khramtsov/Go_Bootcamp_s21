package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "application.h"
#include "window.h"
*/
import "C"
import (
	"unsafe"
)

func main() {
	width, height := C.int(300), C.int(200)
	x, y := C.int(700), C.int(400)

	C.InitApplication()

	ctitle := C.CString("School 21")
	defer C.free(unsafe.Pointer(ctitle))

	window := C.Window_Create(x, y, width, height, ctitle)

	C.Window_MakeKeyAndOrderFront(window)

	C.RunApplication()
}
