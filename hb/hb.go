package hb

// #cgo pkg-config: harfbuzz
// #include <hb.h>
import "C"
import "unsafe"

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func cFeatures(features []Feature) *C.hb_feature_t {
	if len(features) == 0 {
		return nil
	}

	return (*C.hb_feature_t)(unsafe.Pointer(&features[0]))
}
