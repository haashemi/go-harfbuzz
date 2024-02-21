package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Face holds font faces.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-t
type Face *C.hb_face_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-count
func FaceCount(blob Blob) uint32 {
	return uint32(C.hb_face_count(blob))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-create
func FaceCreate(blob Blob, index uint32) Face {
	return C.hb_face_create(blob, C.uint(index))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-create-for-tables
func FaceCreateForTables(referenceTableFunc ReferenceTableFunc, userData unsafe.Pointer, destroy DestroyFunc) Face {
	return C.hb_face_create_for_tables(referenceTableFunc, userData, destroy)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-get-empty
func FaceGetEmpty() Face {
	return C.hb_face_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-reference
func FaceReference(face Face) Face {
	return C.hb_face_reference(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-face.html#hb-face-destroy
func FaceDestroy(face Face) {
	C.hb_face_destroy(face)
}
