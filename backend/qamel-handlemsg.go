package backend

// #include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>
// #include <string.h>
// #include "qamel-handlemsg.h"
import "C"
import (
	"github.com/go-qamel/qamel"
	"unsafe"
)

//export qamelHandleMsgConstructor
func qamelHandleMsgConstructor(ptr unsafe.Pointer) {
	obj := &HandleMsg{}
	obj.Ptr = ptr
	qamel.RegisterObject(ptr, obj)
	obj.init()
}

//export qamelDestroyHandleMsg
func qamelDestroyHandleMsg(ptr unsafe.Pointer) {
	qamel.DeleteObject(ptr)
}

//export qamelHandleMsgMusicControl
func qamelHandleMsgMusicControl(ptr unsafe.Pointer, p0 C.bool, p1 *C.char) {
	obj := qamel.BorrowObject(ptr)
	defer qamel.ReturnObject(ptr)
	if obj == nil {
		return
	}

	objHandleMsg, ok := obj.(*HandleMsg)
	if !ok {
		return
	}

	cgoP0 := bool(p0)
	cgoP1 := C.GoString(p1)
	objHandleMsg.musicControl(cgoP0, cgoP1)
	return
}

// getter and setter

// signals invoker

func (obj *HandleMsg) sendDanMu(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.HandleMsg_SendDanMu(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendGift(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.HandleMsg_SendGift(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendWelCome(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.HandleMsg_SendWelCome(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendWelComeGuard(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.HandleMsg_SendWelComeGuard(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendGreatSailing(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.HandleMsg_SendGreatSailing(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendOnlineChanged(p0 int) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.int(int32(p0))
	C.HandleMsg_SendOnlineChanged(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendFansChanged(p0 int) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.int(int32(p0))
	C.HandleMsg_SendFansChanged(obj.Ptr, cP0)
}

func (obj *HandleMsg) sendMusicURI(p0 string, p1 string, p2 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	cP1 := C.CString(p1)
	defer C.free(unsafe.Pointer(cP1))
	cP2 := C.CString(p2)
	defer C.free(unsafe.Pointer(cP2))
	C.HandleMsg_SendMusicURI(obj.Ptr, cP0, cP1, cP2)
}

// RegisterQmlHandleMsg registers HandleMsg as QML object
func RegisterQmlHandleMsg(uri string, versionMajor int, versionMinor int, qmlName string) {
	cURI := C.CString(uri)
	cQmlName := C.CString(qmlName)
	cVersionMajor := C.int(int32(versionMajor))
	cVersionMinor := C.int(int32(versionMinor))
	defer func() {
		C.free(unsafe.Pointer(cURI))
		C.free(unsafe.Pointer(cQmlName))
	}()

	C.HandleMsg_RegisterQML(cURI, cVersionMajor, cVersionMinor, cQmlName)
}
