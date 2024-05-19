// Code generated by cmd/cgo; DO NOT EDIT.

//line D:\tools\Golang\workspace\WindowsHelperClient\WindowsHelperClient\pkg\ui\cefBrowser.go:1:1
package ui

//#cgo CFLAGS: -I../../cefapi
//#cgo LDFLAGS: -L${SRCDIR}/../../cefapi -L${SRCDIR}/../../Release -lcefapi -lcef
/*
#include <stdio.h>
#include <stdlib.h>

#include "../../cefapi/cefapi.h"

*/
import _ "unsafe"
import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/lxn/win"
)

var mainPageContent = "<html><head><title>Scheme Test(Home Page) From Go</title></head><body bgcolor=\"white\">This contents of this page page are served by the resource handlle , navigate to new page <a href='https://www.sina.com.cn'>Sina</></body></html>"
var offset_ int = 0

func TestMenuItem() {
	createBrowser("aaa", "http://www.sina.com.cn", 0, 0, 0, 0, 0)
}

func InitCef() {
	args := os.Args
	argc :=  /*line :31:10*/_Ctype_int /*line :31:15*/(len(args))
	argv := make([]* /*line :32:18*/_Ctype_char /*line :32:24*/, argc)
	for i, arg := range args {
		argv[i] = ( /*line :34:13*/_Cfunc_CString /*line :34:21*/)(arg)
	}
	( /*line :36:2*/_Cfunc_setBeforePopupCallback /*line :36:25*/)( /*line :36:27*/_Ctype_onBeforePopupFuncProto /*line :36:51*/(( /*line :36:52*/_Cgo_ptr(_Cfpvar_fp_cef_onBeforePopup) /*line :36:70*/)))
	( /*line :37:2*/_Cfunc_setResourceHandlerOpenCallback /*line :37:33*/)( /*line :37:35*/_Ctype_onResourceHandlerOpenFuncProto /*line :37:67*/(( /*line :37:68*/_Cgo_ptr(_Cfpvar_fp_cef_onResourceHandlerOpen) /*line :37:94*/)))
	func() _Ctype_int{ var _cgo0 _Ctype_int = /*line :38:13*/argc; _cgoIndex1 := &/*line :38:46*/argv; _cgo1 := /*line :38:19*/(**_Ctype_char /*line :38:28*/)(unsafe.Pointer(&(*_cgoIndex1)[0])); _cgoCheckPointer(_cgo1, *_cgoIndex1); return /*line :38:56*/_Cfunc_startCef(_cgo0, _cgo1); }()
}

func ShutdownCef() {
	( /*line :42:2*/_Cfunc_shutdownCef /*line :42:14*/)()
}

func createBrowser(title, url string, parent win.HWND, x, y, width, height int) {
	( /*line :46:2*/_Cfunc_createBrowser /*line :46:16*/)(( /*line :46:18*/_Cfunc_CString /*line :46:26*/)(title), ( /*line :46:36*/_Cfunc_CString /*line :46:44*/)(url),  /*line :46:52*/_Ctype_int /*line :46:57*/(int(uintptr(parent))),
		 /*line :47:3*/_Ctype_int /*line :47:8*/(x),  /*line :47:13*/_Ctype_int /*line :47:18*/(y),  /*line :47:23*/_Ctype_int /*line :47:28*/(width),  /*line :47:37*/_Ctype_int /*line :47:42*/(height))
}

func loadUrl(url string) {
	( /*line :51:2*/_Cfunc_loadUrl /*line :51:10*/)(( /*line :51:12*/_Cfunc_CString /*line :51:20*/)(url))
}

func goBack() {
	( /*line :55:2*/_Cfunc_goBack /*line :55:9*/)()
}
func goForward() {
	( /*line :58:2*/_Cfunc_goForward /*line :58:12*/)()
}

func goReload() {
	( /*line :62:2*/_Cfunc_goReload /*line :62:11*/)()
}

//export cef_onBeforePopup
func cef_onBeforePopup(target_url * /*line :66:36*/_Ctype_char /*line :66:42*/)  /*line :66:44*/_Ctype_int /*line :66:49*/ {
	width, height := window.GetSize()
	createBrowser("bbb", ( /*line :68:23*/_Cfunc_GoString /*line :68:32*/)(target_url), GetMainWindowHandler(), 0, toolbarHeight, width, height-toolbarHeight)
	return 1
}

//export cef_onResourceHandlerOpen
func cef_onResourceHandlerOpen(target_url * /*line :73:44*/_Ctype_char /*line :73:50*/, identity  /*line :73:61*/_Ctype_int /*line :73:66*/)  /*line :73:68*/_Ctype_int /*line :73:73*/ {
	fmt.Printf("open resource %s with identity %d\n", ( /*line :74:52*/_Cfunc_GoString /*line :74:61*/)(target_url), identity)
	return 1
}

func setBrowserSize(width, height int) {
	( /*line :79:2*/_Cfunc_setBrowserSize /*line :79:17*/)( /*line :79:19*/_Ctype_int /*line :79:24*/(width),  /*line :79:33*/_Ctype_int /*line :79:38*/(height))
}

//export CopyDataToMemory
func CopyDataToMemory(data * /*line :83:29*/_Ctype_char /*line :83:35*/, size  /*line :83:42*/_Ctype_int /*line :83:47*/) {
	/// 将Go字符串转换为C字符串
	str := "Hello, World!"
	source := []byte(str)
	func() { _cgo0 := /*line :87:17*/unsafe.Pointer(data); _cgoIndex1 := &/*line :87:55*/source; _cgo1 := /*line :87:39*/unsafe.Pointer(&(*_cgoIndex1)[0]); var _cgo2 _Ctype_int = /*line :87:67*/size; _cgoCheckPointer(_cgo0, nil); _cgoCheckPointer(_cgo1, *_cgoIndex1); /*line :87:72*/_Cfunc_goCopyMemory(_cgo0, _cgo1, _cgo2); }()
}

//export cef_onResourceHandlerGetResponseHeaders
func cef_onResourceHandlerGetResponseHeaders(identity  /*line :91:55*/_Ctype_int /*line :91:60*/) (int, string, int) {
	return 200, "text/html", len([]byte(mainPageContent))
}

//export cef_onResourceHandlerRead
func cef_onResourceHandlerRead(identity  /*line :96:41*/_Ctype_int /*line :96:46*/, data_out * /*line :96:58*/_Ctype_char /*line :96:64*/, bytes_to_read int) (int, int) {
	bytes_read := 0
	has_data := 0
	sourceData := []byte(mainPageContent)
	if offset_ < len(sourceData) {
		transfer_size := int(math.Min(float64(bytes_to_read), float64(len(sourceData)-offset_)))
		func() { _cgo0 := /*line :102:18*/unsafe.Pointer(data_out); _cgoIndex1 := &/*line :102:60*/sourceData; _cgo1 := /*line :102:44*/unsafe.Pointer(&(*_cgoIndex1)[offset_]); var _cgo2 _Ctype_int = /*line :102:82*/transfer_size; _cgoCheckPointer(_cgo0, nil); _cgoCheckPointer(_cgo1, *_cgoIndex1); /*line :102:96*/_Cfunc_goCopyMemory(_cgo0, _cgo1, _cgo2); }()
		offset_ = offset_ + transfer_size
		bytes_read = transfer_size
		has_data = 1
	}
	return bytes_read, has_data
}
