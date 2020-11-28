// Code generated by cmd/cgo; DO NOT EDIT.

package govosk

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_VoskModel = _Ctype_struct_VoskModel

type _Ctype_VoskRecognizer = _Ctype_struct_VoskRecognizer

type _Ctype_char int8

type _Ctype_float float32

//go:notinheap
type _Ctype_struct_VoskModel struct{}

//go:notinheap
type _Ctype_struct_VoskRecognizer struct{}

//go:notinheap
type _Ctype_struct_VoskSpkModel struct{}

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_import_static _cgo_50057e25117d_Cfunc_vosk_model_free
//go:linkname __cgofn__cgo_50057e25117d_Cfunc_vosk_model_free _cgo_50057e25117d_Cfunc_vosk_model_free
var __cgofn__cgo_50057e25117d_Cfunc_vosk_model_free byte
var _cgo_50057e25117d_Cfunc_vosk_model_free = unsafe.Pointer(&__cgofn__cgo_50057e25117d_Cfunc_vosk_model_free)

//go:cgo_unsafe_args
func _Cfunc_vosk_model_free(p0 *_Ctype_struct_VoskModel) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_50057e25117d_Cfunc_vosk_model_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_50057e25117d_Cfunc_vosk_model_new
//go:linkname __cgofn__cgo_50057e25117d_Cfunc_vosk_model_new _cgo_50057e25117d_Cfunc_vosk_model_new
var __cgofn__cgo_50057e25117d_Cfunc_vosk_model_new byte
var _cgo_50057e25117d_Cfunc_vosk_model_new = unsafe.Pointer(&__cgofn__cgo_50057e25117d_Cfunc_vosk_model_new)

//go:cgo_unsafe_args
func _Cfunc_vosk_model_new(p0 *_Ctype_char) (r1 *_Ctype_struct_VoskModel) {
	_cgo_runtime_cgocall(_cgo_50057e25117d_Cfunc_vosk_model_new, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_50057e25117d_Cfunc_vosk_recognizer_free
//go:linkname __cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_free _cgo_50057e25117d_Cfunc_vosk_recognizer_free
var __cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_free byte
var _cgo_50057e25117d_Cfunc_vosk_recognizer_free = unsafe.Pointer(&__cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_free)

//go:cgo_unsafe_args
func _Cfunc_vosk_recognizer_free(p0 *_Ctype_struct_VoskRecognizer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_50057e25117d_Cfunc_vosk_recognizer_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_50057e25117d_Cfunc_vosk_recognizer_new
//go:linkname __cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_new _cgo_50057e25117d_Cfunc_vosk_recognizer_new
var __cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_new byte
var _cgo_50057e25117d_Cfunc_vosk_recognizer_new = unsafe.Pointer(&__cgofn__cgo_50057e25117d_Cfunc_vosk_recognizer_new)

//go:cgo_unsafe_args
func _Cfunc_vosk_recognizer_new(p0 *_Ctype_struct_VoskModel, p1 _Ctype_float) (r1 *_Ctype_struct_VoskRecognizer) {
	_cgo_runtime_cgocall(_cgo_50057e25117d_Cfunc_vosk_recognizer_new, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}