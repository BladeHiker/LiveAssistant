package backend

// #include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>
// #include <string.h>
// #include "qamel-connectfeedback.h"
import "C"
import (
	"github.com/go-qamel/qamel"
	"unsafe"
)

//export qamelConnectFeedBackConstructor
func qamelConnectFeedBackConstructor(ptr unsafe.Pointer) {
	obj := &ConnectFeedBack{}
	obj.Ptr = ptr
	qamel.RegisterObject(ptr, obj)
}

//export qamelDestroyConnectFeedBack
func qamelDestroyConnectFeedBack(ptr unsafe.Pointer) {
	qamel.DeleteObject(ptr)
}

//export qamelConnectFeedBackReceiveRoomID
func qamelConnectFeedBackReceiveRoomID(ptr unsafe.Pointer, p0 C.int) {
	obj := qamel.BorrowObject(ptr)
	defer qamel.ReturnObject(ptr)
	if obj == nil {
		return
	}

	objConnectFeedBack, ok := obj.(*ConnectFeedBack)
	if !ok {
		return
	}

	cgoP0 := int(int32(p0))
	objConnectFeedBack.receiveRoomID(cgoP0)
	return
}

// getter and setter

// signals invoker

func (obj *ConnectFeedBack) sendFansNums(p0 int) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.int(int32(p0))
	C.ConnectFeedBack_SendFansNums(obj.Ptr, cP0)
}

func (obj *ConnectFeedBack) sendCompInfo(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.ConnectFeedBack_SendCompInfo(obj.Ptr, cP0)
}

func (obj *ConnectFeedBack) sendErr(p0 int) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.int(int32(p0))
	C.ConnectFeedBack_SendErr(obj.Ptr, cP0)
}

// RegisterQmlConnectFeedBack registers ConnectFeedBack as QML object
func RegisterQmlConnectFeedBack(uri string, versionMajor int, versionMinor int, qmlName string) {
	cURI := C.CString(uri)
	cQmlName := C.CString(qmlName)
	cVersionMajor := C.int(int32(versionMajor))
	cVersionMinor := C.int(int32(versionMinor))
	defer func() {
		C.free(unsafe.Pointer(cURI))
		C.free(unsafe.Pointer(cQmlName))
	}()

	C.ConnectFeedBack_RegisterQML(cURI, cVersionMajor, cVersionMinor, cQmlName)
}
