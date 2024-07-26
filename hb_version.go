package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

const (
	HB_VERSION_MAJOR  = C.HB_VERSION_MAJOR  // The major component of the library version available at compile-time.
	HB_VERSION_MICRO  = C.HB_VERSION_MICRO  // The micro component of the library version available at compile-time.
	HB_VERSION_MINOR  = C.HB_VERSION_MINOR  // The minor component of the library version available at compile-time.
	HB_VERSION_STRING = C.HB_VERSION_STRING // A string literal containing the library version available at compile-time.
)

// Version returns library version as three uint32 components.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-version.html#hb-version
func Version() (major, minor, micro uint32) {
	C.hb_version((*C.uint)(unsafe.Pointer(&major)), (*C.uint)(unsafe.Pointer(&minor)), (*C.uint)(unsafe.Pointer(&micro)))
	return
}

// VersionAtLeast tests the library version against a minimum value, as three
// integer components.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-version.html#hb-version-atleast
func VersionAtLeast(major, minor, micro uint32) bool {
	return C.hb_version_atleast(C.uint(major), C.uint(minor), C.uint(micro)) == 1
}

// VersionString returns library version as a string with three components.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-version.html#hb-version-string
func VersionString() string {
	v := C.hb_version_string()
	return C.GoString(v)
}
