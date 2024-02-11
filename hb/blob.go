package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

type Blob *C.hb_blob_t

type MemoryMode C.hb_memory_mode_t

const (
	MemoryModeDuplicate               MemoryMode = C.HB_MEMORY_MODE_DUPLICATE
	MemoryModeReadonly                MemoryMode = C.HB_MEMORY_MODE_READONLY
	MemoryModeWritable                MemoryMode = C.HB_MEMORY_MODE_WRITABLE
	MemoryModeReadonlyMayMakeWritable MemoryMode = C.HB_MEMORY_MODE_READONLY_MAY_MAKE_WRITABLE
)

// TODO: support user_data & its destroyer.
func BlobCreate(data string, length int, mode MemoryMode) Blob {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return C.hb_blob_create(cData, C.uint(length), C.hb_memory_mode_t(mode), nil, nil)
}

// TODO: support user_data & its destroyer.
func BlobCreateOrFail(data string, length int, mode MemoryMode) Blob {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return C.hb_blob_create_or_fail(cData, C.uint(length), C.hb_memory_mode_t(mode), nil, nil)
}

func BlobCreateFromFile(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file(file_name)
}

func BlobCreateFromFileOrFail(filename string) Blob {
	file_name := C.CString(filename)
	defer C.free(unsafe.Pointer(file_name))

	return C.hb_blob_create_from_file_or_fail(file_name)
}

func BlobCreateSubBlob(parent Blob, offset, length uint) Blob {
	return C.hb_blob_create_sub_blob(parent, C.uint(offset), C.uint(length))
}

func BlobCopyWritableOrFail(blob Blob) Blob {
	return C.hb_blob_copy_writable_or_fail(blob)
}

func BlobGetEmpty() Blob {
	return C.hb_blob_get_empty()
}

func BlobReference(blob Blob) Blob {
	return C.hb_blob_reference(blob)
}

func BlobDestroy(blob Blob) {
	C.hb_blob_destroy(blob)
}

// TODO: C.hb_blob_set_user_data
// TODO: C.hb_blob_get_user_data

func BlobMakeImmutable(blob Blob) {
	C.hb_blob_make_immutable(blob)
}

func BlobIsImmutable(blob Blob) bool {
	return C.hb_blob_is_immutable(blob) == 1
}

func BlobGetData(blob Blob) string {
	var length C.uint
	data := C.hb_blob_get_data(blob, &length)

	return C.GoStringN(data, C.int(length))
}

func BlobGetLength(blob Blob) uint {
	return uint(C.hb_blob_get_length(blob))
}

// Developer notes:
// There's no need to implement hb_blob_get_data_writable as we're already copying the data.
