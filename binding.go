package language

// #include "binding.h"
import "C"
import "unsafe"

// detectLanguage calls the C function
func detectLanguage(text string, plain bool) *Language {
	length := C.int(len(text))
	buffer := C.CString(text)
	defer C.free(unsafe.Pointer(buffer))
	result := C.struct_language{}
	C.detect_language(buffer, length, C._Bool(plain), &result)
	return &Language{
		Code:     C.GoString(result.code),
		Name:     C.GoString(result.name),
		Reliable: result.reliable != 0,
	}
}
