package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import (
	"unsafe"
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape.html#hb-shape
func Shape(font Font, buffer Buffer, features []Feature) {
	C.hb_shape(font, buffer, cFeatures(features), C.uint(len(features)))
}

// TODO: Support shapers as argument
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape.html#hb-shape-full
func ShapeFull(font Font, buffer Buffer, features []Feature) {
	C.hb_shape_full(font, buffer, (*C.hb_feature_t)(unsafe.Pointer(&features[0])), C.uint(len(features)), C.hb_shape_list_shapers())
}

// ShapeListShapers returns the list of shapers supported by HarfBuzz.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape.html#hb-shape-list-shapers
func ShapeListShapers() (shapers []string) {
	res := C.hb_shape_list_shapers()

	for data := *res; data != nil; data = *res {
		shapers = append(shapers, C.GoString(data))

		res = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(res)) + unsafe.Sizeof(res)))
	}

	return shapers
}
