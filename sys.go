/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
*/
import "C"
import (
	"unsafe"
)

//PySys_GetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_GetObject
func PySys_GetObject(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PySys_GetObject(cname))
}

//PySys_SetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_SetObject
func PySys_SetObject(name string, v *PyObject) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PySys_SetObject(cname, toc(v)))
}

//PySys_ResetWarnOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_ResetWarnOptions
func PySys_ResetWarnOptions() {
	C.PySys_ResetWarnOptions()
}

//PySys_GetXOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_GetXOptions
func PySys_GetXOptions() *PyObject {
	return togo(C.PySys_GetXOptions())
}
