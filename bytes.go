/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

// Bytes : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Type
var Bytes = togo((*C.PyObject)(unsafe.Pointer(&C.PyBytes_Type)))

// PyBytes_Check : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Check
func PyBytes_Check(o *PyObject) bool {
	return C._go_PyBytes_Check(toc(o)) != 0
}

// PyBytes_CheckExact : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_CheckExact
func PyBytes_CheckExact(o *PyObject) bool {
	return C._go_PyBytes_CheckExact(toc(o)) != 0
}

// PyBytes_FromString : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromString
func PyBytes_FromString(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyBytes_FromString(cstr))
}

// PyBytes_FromStringAndSize : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromStringAndSize
func PyBytes_FromStringAndSize(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyBytes_FromStringAndSize(cstr, C.Py_ssize_t(len(str))))
}

// PyBytes_FromObject : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromObject
func PyBytes_FromObject(o *PyObject) *PyObject {
	return togo(C.PyBytes_FromObject(toc(o)))
}

// PyBytes_Size : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Size
func PyBytes_Size(o *PyObject) int {
	return int(C.PyBytes_Size(toc(o)))
}

// PyBytes_AsString : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_AsString
func PyBytes_AsString(o *PyObject) string {
	return C.GoStringN(C.PyBytes_AsString(toc(o)), C.int(C.PyBytes_Size(toc(o))))
}

// PyBytes_Concat : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_Concat
func PyBytes_Concat(bytes, newpart *PyObject) *PyObject {
	cbytes := toc(bytes)
	C.PyBytes_Concat(&cbytes, toc(newpart))
	return togo(cbytes)
}

// PyBytes_ConcatAndDel : https://docs.python.org/3/c-api/bytes.html#c.PyBytes_ConcatAndDel
func PyBytes_ConcatAndDel(bytes, newpart *PyObject) *PyObject {
	cbytes := toc(bytes)
	C.PyBytes_ConcatAndDel(&cbytes, toc(newpart))
	return togo(cbytes)
}

// PyBytes_FromByteSlice uses https://docs.python.org/3/c-api/bytes.html#c.PyBytes_FromStringAndSize but with []byte
func PyBytes_FromByteSlice(bytes []byte) *PyObject {
	pbytes := C.CBytes(bytes)
	defer C.free(pbytes)

	cstr := (*C.char)(pbytes)

	return togo(C.PyBytes_FromStringAndSize(cstr, C.Py_ssize_t(len(bytes))))
}

// PyBytes_AsByteSlice is equivalent to PyBytes_AsString but returns byte slices
func PyBytes_AsByteSlice(o *PyObject) []byte {
	cstr := C.PyBytes_AsString(toc(o))
	size := C.PyBytes_Size(toc(o))

	return C.GoBytes(unsafe.Pointer(cstr), C.int(size))
}
