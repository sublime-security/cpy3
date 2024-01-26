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
	"fmt"
	"unsafe"
)

// Py_Initialize : https://docs.python.org/3/c-api/init.html#c.Py_Initialize
func Py_Initialize() {
	C.Py_Initialize()
}

// Py_InitializeEx : https://docs.python.org/3/c-api/init.html#c.Py_InitializeEx
func Py_InitializeEx(initsigs bool) {
	if initsigs {
		C.Py_InitializeEx(1)
	} else {
		C.Py_InitializeEx(0)
	}
}

// Py_IsInitialized : https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func Py_IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

// Py_FinalizeEx : https://docs.python.org/3/c-api/init.html#c.Py_FinalizeEx
func Py_FinalizeEx() int {
	return int(C.Py_FinalizeEx())
}

// Py_Finalize : https://docs.python.org/3/c-api/init.html#c.Py_Finalize
func Py_Finalize() {
	C.Py_Finalize()
}

// Py_GetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func Py_GetProgramName() (string, error) {
	wcname := C.Py_GetProgramName()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func Py_GetPrefix() (string, error) {
	wcname := C.Py_GetPrefix()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetExecPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() (string, error) {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetProgramFullPath : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() (string, error) {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetPath : https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() (string, error) {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

// Py_GetVersion : https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func Py_GetVersion() string {
	cversion := C.Py_GetVersion()
	return C.GoString(cversion)
}

// Py_GetPlatform : https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func Py_GetPlatform() string {
	cplatform := C.Py_GetPlatform()
	return C.GoString(cplatform)
}

// Py_GetCopyright : https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func Py_GetCopyright() string {
	ccopyright := C.Py_GetCopyright()
	return C.GoString(ccopyright)
}

// Py_GetCompiler : https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func Py_GetCompiler() string {
	ccompiler := C.Py_GetCompiler()
	return C.GoString(ccompiler)
}

// Py_GetBuildInfo : https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func Py_GetBuildInfo() string {
	cbuildInfo := C.Py_GetBuildInfo()
	return C.GoString(cbuildInfo)
}

// Py_GetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() (string, error) {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return "", nil
	}
	chome := C.Py_EncodeLocale(wchome, nil)
	if chome == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome), nil
}
