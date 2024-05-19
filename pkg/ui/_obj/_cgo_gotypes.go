// Code generated by cmd/cgo; DO NOT EDIT.

package ui

import "unsafe"

import "syscall"

import _cgopackage "runtime/cgo"

type _ _cgopackage.Incomplete
var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
//go:linkname _Cgo_no_callback runtime.cgoNoCallback
func _Cgo_no_callback(bool)
type _Ctype__GoString_ string

type _Ctype_char int8

type _Ctype_int int32

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_longlong int64

type _Ctype_onBeforePopupFuncProto *[0]byte

type _Ctype_onResourceHandlerOpenFuncProto *[0]byte

type _Ctype_ptrdiff_t = _Ctype_longlong

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
//go:noescape
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
//go:noescape
func _cgoCheckResult(interface{})
//go:linkname __cgo_cef_onBeforePopup cef_onBeforePopup
//go:cgo_import_static cef_onBeforePopup
var __cgo_cef_onBeforePopup byte
var _Cfpvar_fp_cef_onBeforePopup unsafe.Pointer = (unsafe.Pointer)(unsafe.Pointer(&__cgo_cef_onBeforePopup))
//go:linkname __cgo_cef_onResourceHandlerOpen cef_onResourceHandlerOpen
//go:cgo_import_static cef_onResourceHandlerOpen
var __cgo_cef_onResourceHandlerOpen byte
var _Cfpvar_fp_cef_onResourceHandlerOpen unsafe.Pointer = (unsafe.Pointer)(unsafe.Pointer(&__cgo_cef_onResourceHandlerOpen))


// CString converts the Go string s to a C string.
//
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func _Cfunc_CString(s string) *_Ctype_char {
	if len(s)+1 <= 0 {
		panic("string too large")
	}
	p := _cgo_cmalloc(uint64(len(s)+1))
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, len(s)+1, len(s)+1}
	b := *(*[]byte)(unsafe.Pointer(&sliceHeader))
	copy(b, s)
	b[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

// GoString converts the C string p into a Go string.
func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_createBrowser
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_createBrowser _cgo_8c88dd3c38f3_Cfunc_createBrowser
var __cgofn__cgo_8c88dd3c38f3_Cfunc_createBrowser byte
var _cgo_8c88dd3c38f3_Cfunc_createBrowser = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_createBrowser)

//go:cgo_unsafe_args
func _Cfunc_createBrowser(p0 *_Ctype_char, p1 *_Ctype_char, p2 _Ctype_int, p3 _Ctype_int, p4 _Ctype_int, p5 _Ctype_int, p6 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_createBrowser, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
		_Cgo_use(p4)
		_Cgo_use(p5)
		_Cgo_use(p6)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_goBack
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_goBack _cgo_8c88dd3c38f3_Cfunc_goBack
var __cgofn__cgo_8c88dd3c38f3_Cfunc_goBack byte
var _cgo_8c88dd3c38f3_Cfunc_goBack = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_goBack)

//go:cgo_unsafe_args
func _Cfunc_goBack() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_goBack, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_goCopyMemory
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_goCopyMemory _cgo_8c88dd3c38f3_Cfunc_goCopyMemory
var __cgofn__cgo_8c88dd3c38f3_Cfunc_goCopyMemory byte
var _cgo_8c88dd3c38f3_Cfunc_goCopyMemory = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_goCopyMemory)

//go:cgo_unsafe_args
func _Cfunc_goCopyMemory(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 _Ctype_int) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_goCopyMemory, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_goForward
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_goForward _cgo_8c88dd3c38f3_Cfunc_goForward
var __cgofn__cgo_8c88dd3c38f3_Cfunc_goForward byte
var _cgo_8c88dd3c38f3_Cfunc_goForward = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_goForward)

//go:cgo_unsafe_args
func _Cfunc_goForward() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_goForward, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_goReload
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_goReload _cgo_8c88dd3c38f3_Cfunc_goReload
var __cgofn__cgo_8c88dd3c38f3_Cfunc_goReload byte
var _cgo_8c88dd3c38f3_Cfunc_goReload = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_goReload)

//go:cgo_unsafe_args
func _Cfunc_goReload() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_goReload, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_loadUrl
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_loadUrl _cgo_8c88dd3c38f3_Cfunc_loadUrl
var __cgofn__cgo_8c88dd3c38f3_Cfunc_loadUrl byte
var _cgo_8c88dd3c38f3_Cfunc_loadUrl = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_loadUrl)

//go:cgo_unsafe_args
func _Cfunc_loadUrl(p0 *_Ctype_char) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_loadUrl, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback _cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback
var __cgofn__cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback byte
var _cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback)

//go:cgo_unsafe_args
func _Cfunc_setBeforePopupCallback(p0 *[0]byte) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_setBeforePopupCallback, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_setBrowserSize
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_setBrowserSize _cgo_8c88dd3c38f3_Cfunc_setBrowserSize
var __cgofn__cgo_8c88dd3c38f3_Cfunc_setBrowserSize byte
var _cgo_8c88dd3c38f3_Cfunc_setBrowserSize = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_setBrowserSize)

//go:cgo_unsafe_args
func _Cfunc_setBrowserSize(p0 _Ctype_int, p1 _Ctype_int) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_setBrowserSize, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback _cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback
var __cgofn__cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback byte
var _cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback)

//go:cgo_unsafe_args
func _Cfunc_setResourceHandlerOpenCallback(p0 *[0]byte) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_setResourceHandlerOpenCallback, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_shutdownCef
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_shutdownCef _cgo_8c88dd3c38f3_Cfunc_shutdownCef
var __cgofn__cgo_8c88dd3c38f3_Cfunc_shutdownCef byte
var _cgo_8c88dd3c38f3_Cfunc_shutdownCef = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_shutdownCef)

//go:cgo_unsafe_args
func _Cfunc_shutdownCef() (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_shutdownCef, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc_startCef
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc_startCef _cgo_8c88dd3c38f3_Cfunc_startCef
var __cgofn__cgo_8c88dd3c38f3_Cfunc_startCef byte
var _cgo_8c88dd3c38f3_Cfunc_startCef = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc_startCef)

//go:cgo_unsafe_args
func _Cfunc_startCef(p0 _Ctype_int, p1 **_Ctype_char) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc_startCef, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_export_dynamic cef_onBeforePopup
//go:linkname _cgoexp_8c88dd3c38f3_cef_onBeforePopup _cgoexp_8c88dd3c38f3_cef_onBeforePopup
//go:cgo_export_static _cgoexp_8c88dd3c38f3_cef_onBeforePopup
func _cgoexp_8c88dd3c38f3_cef_onBeforePopup(a *struct {
		p0 *_Ctype_char
		r0 _Ctype_int
	}) {
	a.r0 = cef_onBeforePopup(a.p0)
}
//go:cgo_export_dynamic cef_onResourceHandlerOpen
//go:linkname _cgoexp_8c88dd3c38f3_cef_onResourceHandlerOpen _cgoexp_8c88dd3c38f3_cef_onResourceHandlerOpen
//go:cgo_export_static _cgoexp_8c88dd3c38f3_cef_onResourceHandlerOpen
func _cgoexp_8c88dd3c38f3_cef_onResourceHandlerOpen(a *struct {
		p0 *_Ctype_char
		p1 _Ctype_int
		r0 _Ctype_int
	}) {
	a.r0 = cef_onResourceHandlerOpen(a.p0, a.p1)
}
//go:cgo_export_dynamic CopyDataToMemory
//go:linkname _cgoexp_8c88dd3c38f3_CopyDataToMemory _cgoexp_8c88dd3c38f3_CopyDataToMemory
//go:cgo_export_static _cgoexp_8c88dd3c38f3_CopyDataToMemory
func _cgoexp_8c88dd3c38f3_CopyDataToMemory(a *struct {
		p0 *_Ctype_char
		p1 _Ctype_int
	}) {
	CopyDataToMemory(a.p0, a.p1)
}
//go:cgo_export_dynamic cef_onResourceHandlerGetResponseHeaders
//go:linkname _cgoexp_8c88dd3c38f3_cef_onResourceHandlerGetResponseHeaders _cgoexp_8c88dd3c38f3_cef_onResourceHandlerGetResponseHeaders
//go:cgo_export_static _cgoexp_8c88dd3c38f3_cef_onResourceHandlerGetResponseHeaders
func _cgoexp_8c88dd3c38f3_cef_onResourceHandlerGetResponseHeaders(a *struct {
		p0 _Ctype_int
		r0 int
		r1 string
		r2 int
	}) {
	a.r0, a.r1, a.r2 = cef_onResourceHandlerGetResponseHeaders(a.p0)
	_cgoCheckResult(a.r1)
}
//go:cgo_export_dynamic cef_onResourceHandlerRead
//go:linkname _cgoexp_8c88dd3c38f3_cef_onResourceHandlerRead _cgoexp_8c88dd3c38f3_cef_onResourceHandlerRead
//go:cgo_export_static _cgoexp_8c88dd3c38f3_cef_onResourceHandlerRead
func _cgoexp_8c88dd3c38f3_cef_onResourceHandlerRead(a *struct {
		p0 _Ctype_int
		p1 *_Ctype_char
		p2 int
		r0 int
		r1 int
	}) {
	a.r0, a.r1 = cef_onResourceHandlerRead(a.p0, a.p1, a.p2)
}

//go:cgo_import_static _cgo_8c88dd3c38f3_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_8c88dd3c38f3_Cfunc__Cmalloc _cgo_8c88dd3c38f3_Cfunc__Cmalloc
var __cgofn__cgo_8c88dd3c38f3_Cfunc__Cmalloc byte
var _cgo_8c88dd3c38f3_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_8c88dd3c38f3_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_8c88dd3c38f3_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}