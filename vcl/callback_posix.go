// +build linux,cgo darwin,cgo

package vcl

// extern void* doEventCallbackProc(void* f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(void* f, void* msg, void* handled);
// static void* doGetMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	callbackProc(uintptr(f), uintptr(args), int(argcount))
	return unsafe.Pointer(uintptr(0))
}

//export doMessageCallbackProc
func doMessageCallbackProc(f unsafe.Pointer, msg, handled unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(uintptr(f), uintptr(msg), int(handled))
	return unsafe.Pointer(uintptr(0))
}

var (
	eventCallback   = uintptr(C.doGetEventCallbackAddr())
	messageCallback = uintptr(C.doGetMessageCallbackAddr())
)
