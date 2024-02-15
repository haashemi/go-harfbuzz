package hb

// #cgo pkg-config: harfbuzz
import "C"

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
