package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Blob wraps a chunk of binary data and facilitates its lifecycle management
// between a client program and HarfBuzz.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-t
type Blob *C.hb_blob_t

// MemoryMode holds the memory modes available to client programs.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-memory-mode-t
type MemoryMode C.hb_memory_mode_t

const (
	// HarfBuzz immediately makes a copy of the data.
	//
	// In no case shall the HarfBuzz client modify memory that is passed to
	// HarfBuzz in a blob. If there is any such possibility, this mode should be
	// used such that HarfBuzz makes a copy immediately,
	MemoryModeDuplicate MemoryMode = C.HB_MEMORY_MODE_DUPLICATE

	// HarfBuzz client will never modify the data, and HarfBuzz will never modify
	// the data.
	//
	// Use this if it's ok for Harfbuzz client to modify memory that is passed
	// too Harfbuzz in a blob, unless you really really really know what you are
	// doing.
	MemoryModeReadonly MemoryMode = C.HB_MEMORY_MODE_READONLY

	// HarfBuzz client made a copy f the data solely for HarfBuzz, so HarfBuzz
	// may modify the data.
	//
	// This mode is appropriate if you really made a copy of data solely for the
	// purpose of passing to HarfBuzz and doing that just once (no reuse!).
	MemoryModeWritable MemoryMode = C.HB_MEMORY_MODE_WRITABLE

	MemoryModeReadonlyMayMakeWritable MemoryMode = C.HB_MEMORY_MODE_READONLY_MAY_MAKE_WRITABLE
)

// BlobCreate creates a new Blob wrapping data. The mode parameter is used to
// negotiate ownership and lifecycle of data.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create
//
// TODO: support user_data & its destroyer.
// TODO: research on the blob data memory management.
func BlobCreate(data string, length int, mode MemoryMode) Blob {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return C.hb_blob_create(cData, C.uint(length), C.hb_memory_mode_t(mode), nil, nil)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-or-fail
//
// TODO: support user_data & its destroyer.
func BlobCreateOrFail(data string, length int, mode MemoryMode) Blob {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return C.hb_blob_create_or_fail(cData, C.uint(length), C.hb_memory_mode_t(mode), nil, nil)
}

// BlobCreateFromFile creates a new blob containing the data from the specified
// binary font file.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-from-file
func BlobCreateFromFile(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file(file_name)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-from-file-or-fail
func BlobCreateFromFileOrFail(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file_or_fail(file_name)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-create-sub-blob
func BlobCreateSubBlob(parent Blob, offset, length uint) Blob {
	return C.hb_blob_create_sub_blob(parent, C.uint(offset), C.uint(length))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-copy-writable-or-fail
func BlobCopyWritableOrFail(blob Blob) Blob {
	return C.hb_blob_copy_writable_or_fail(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-empty
func BlobGetEmpty() Blob {
	return C.hb_blob_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-reference
func BlobReference(blob Blob) Blob {
	return C.hb_blob_reference(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-destroy
func BlobDestroy(blob Blob) {
	C.hb_blob_destroy(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-set-user-data
// TODO: C.hb_blob_set_user_data

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-user-data
// TODO: C.hb_blob_get_user_data

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-make-immutable
func BlobMakeImmutable(blob Blob) {
	C.hb_blob_make_immutable(blob)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-is-immutable
func BlobIsImmutable(blob Blob) bool {
	return C.hb_blob_is_immutable(blob) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-data
func BlobGetData(blob Blob) string {
	var length C.uint
	data := C.hb_blob_get_data(blob, &length)

	return C.GoStringN(data, C.int(length))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-blob.html#hb-blob-get-length
func BlobGetLength(blob Blob) uint {
	return uint(C.hb_blob_get_length(blob))
}

// Developer notes:
// There's no need to implement hb_blob_get_data_writable as we're already copying the data. (AFAIK)
