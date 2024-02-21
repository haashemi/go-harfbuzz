package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"

import (
	"unsafe"
)

// Buffer is the main structure holding the input text and its properties
// before shaping, and output glyphs and their information after shaping.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-t
type Buffer *C.hb_buffer_t

// TODO: Document + Find a better for conversion between C & Go.
type GlyphInfo struct {
	CodePoint uint32
	Mask      uint32
	Cluster   uint32
	Var1      [4]byte
	Var2      [4]byte
}

// TODO: Document + Find a better for conversion between C & Go.
type GlyphPosition struct {
	XAdvance uint32
	YAdvance uint32
	XOffset  uint32
	YOffset  uint32
	Var      [4]byte
}

// ContentType is type of buffer's contents.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-content-type-t
type ContentType C.hb_buffer_content_type_t

const (
	ContentTypeINVALID ContentType = C.HB_BUFFER_CONTENT_TYPE_INVALID // Initial value for new buffer.
	ContentTypeUNICODE ContentType = C.HB_BUFFER_CONTENT_TYPE_UNICODE // The buffer contains input characters (before shaping).
	ContentTypeGLYPHS  ContentType = C.HB_BUFFER_CONTENT_TYPE_GLYPHS  // The buffer contains output glyphs (after shaping).
)

// BufferCreate returns a newly allocated Buffer. This function never returns nil.
// Close the Buffer using BufferDestroy.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-create
func BufferCreate() Buffer {
	return C.hb_buffer_create()
}

// BufferAllocationSuccessful Checks if allocating memory for the buffer succeeded.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-allocation-successful
func BufferAllocationSuccessful(buffer Buffer) bool {
	return C.hb_buffer_allocation_successful(buffer) == 1
}

// BufferCreateSimilar returns a newly allocated Buffer which is configured
// similarly to src.
// Close the Buffer using BufferDestroy.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-create-similar
func BufferCreateSimilar(src Buffer) Buffer {
	return C.hb_buffer_create_similar(src)
}

// BufferGetEmpty returns an empty Buffer.
// Close the Buffer using BufferDestroy.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-empty
func BufferGetEmpty() Buffer {
	return C.hb_buffer_get_empty()
}

// BufferReference increases the reference count on the buffer by one.
// Decrease the reference count using BufferDestroy.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-reference
func BufferReference(buffer Buffer) Buffer {
	return C.hb_buffer_reference(buffer)
}

// BufferDestroy De-allocates the buffer. It decreases the reference counts by one.
// Once reference counts reach zero, then the buffer and all associated resources
// are freed.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-destroy
func BufferDestroy(buffer Buffer) {
	C.hb_buffer_destroy(buffer)
}

func BufferSetUserData(buffer Buffer, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_buffer_set_user_data(buffer, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

func BufferGetUserData(buffer Buffer, key *UserDataKey) unsafe.Pointer {
	return C.hb_buffer_get_user_data(buffer, (*C.hb_user_data_key_t)(key))
}

// BufferReset resets the buffer to its initial status, as if it was just newly
// created with BufferCreate.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-reset
func BufferReset(buffer Buffer) {
	C.hb_buffer_reset(buffer)
}

// BufferClearContents resets the buffer to its initial status, as if it was just
// newly created with BufferCreate, but does not clear the Unicode functions and
// the replacement code point.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-clear-contents
func BufferClearContents(buffer Buffer) {
	C.hb_buffer_clear_contents(buffer)
}

// BufferPreAllocate Pre allocates memory for buffer to fit at least size number
// of items. returns true on successful allocation.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-pre-allocate
func BufferPreAllocate(buffer Buffer, size uint32) bool {
	return C.hb_buffer_pre_allocate(buffer, C.uint(size)) == 1
}

// TODO: C.hb_buffer_add
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add

// TODO: C.hb_buffer_add_codepoints
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add-codepoints

// TODO: C.hb_buffer_add_utf32
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add-utf32

// TODO: C.hb_buffer_add_utf16
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add-utf16

// TODO: Write a proper description for the functions bellow.

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add-utf8
func BufferAddUTF8(buffer Buffer, text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.hb_buffer_add_utf8(buffer, cText, -1, 0, -1)
}

// TODO: C.hb_buffer_add_latin1
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-add-latin1

// BufferAppend appends part of the src buffer to the dst buffer.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-append
func BufferAppend(dst, src Buffer, start, end uint32) {
	C.hb_buffer_append(dst, src, C.uint(start), C.uint(end))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-content-type
func BufferSetContentType(buffer Buffer, contentType ContentType) {
	C.hb_buffer_set_content_type(buffer, C.hb_buffer_content_type_t(contentType))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-content-type
func BufferGetContentType(buffer Buffer) ContentType {
	return ContentType(C.hb_buffer_get_content_type(buffer))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-direction
func BufferSetDirection(buffer Buffer, direction Direction) {
	C.hb_buffer_set_direction(buffer, C.hb_direction_t(direction))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-direction
func BufferGetDirection(buffer Buffer) Direction {
	return Direction(C.hb_buffer_get_direction(buffer))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-script
func BufferSetScript(buffer Buffer, script Script) {
	C.hb_buffer_set_script(buffer, C.hb_script_t(script))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-script
func BufferGetScript(buffer Buffer) Script {
	return Script(C.hb_buffer_get_script(buffer))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-language
// TODO: C.hb_buffer_set_language

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-language
// TODO: C.hb_buffer_get_language

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-flags
// TODO: C.hb_buffer_set_flags

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-flags
// TODO: C.hb_buffer_get_flags

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-cluster-level
// TODO: C.hb_buffer_set_cluster_level

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-cluster-level
// TODO: C.hb_buffer_get_cluster_level

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-length
// TODO: C.hb_buffer_set_length

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-length
// TODO: C.hb_buffer_get_length

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-segment-properties
// TODO: C.hb_buffer_set_segment_properties

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-segment-properties
// TODO: C.hb_buffer_get_segment_properties

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-guess-segment-properties
func BufferGuessSegmentProperties(buffer Buffer) {
	C.hb_buffer_guess_segment_properties(buffer)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-unicode-funcs
// TODO: C.hb_buffer_set_unicode_funcs

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-unicode-funcs
// TODO: C.hb_buffer_get_unicode_funcs

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-glyph-infos
func BufferGetGlyphInfos(buffer Buffer, length *uint32) (res []GlyphInfo) {
	data := C.hb_buffer_get_glyph_infos(buffer, (*C.uint)(unsafe.Pointer(length)))
	defer C.free(unsafe.Pointer(data))

	size := unsafe.Sizeof(C.hb_glyph_info_t{})

	for i := uint32(0); i < *length; i++ {
		gi := *(*GlyphInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + size*uintptr(i)))
		res = append(res, gi)
	}

	return res
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-glyph-info-get-glyph-flags
// TODO: C.hb_glyph_info_get_glyph_flags

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-glyph-positions
func BufferGetGlyphPositions(buffer Buffer, length *uint32) (res []GlyphPosition) {
	data := C.hb_buffer_get_glyph_positions(buffer, (*C.uint)(unsafe.Pointer(length)))
	defer C.free(unsafe.Pointer(data))

	size := unsafe.Sizeof(C.hb_glyph_position_t{})

	for i := uint32(0); i < *length; i++ {
		gp := *(*GlyphPosition)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + size*uintptr(i)))
		res = append(res, gp)
	}

	return res
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-has-positions
// TODO: C.hb_buffer_has_positions

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-invisible-glyph
// TODO: C.hb_buffer_set_invisible_glyph

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-invisible-glyph
// TODO: C.hb_buffer_get_invisible_glyph

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-not-found-glyph
// TODO: C.hb_buffer_set_not_found_glyph

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-not-found-glyph
// TODO: C.hb_buffer_get_not_found_glyph

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-replacement-codepoint
// TODO: C.hb_buffer_set_replacement_codepoint

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-replacement-codepoint
// TODO: C.hb_buffer_get_replacement_codepoint

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-normalize-glyphs
// TODO: C.hb_buffer_normalize_glyphs

// BufferReverse reverses the buffer contents.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-reverse
func BufferReverse(buffer Buffer) {
	C.hb_buffer_reverse(buffer)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-reverse-range
// TODO: C.hb_buffer_reverse_range

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-reverse-clusters
// TODO: C.hb_buffer_reverse_clusters

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize
// TODO: C.hb_buffer_serialize

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-glyphs
// TODO: C.hb_buffer_serialize_glyphs

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-deserialize-glyphs
// TODO: C.hb_buffer_deserialize_glyphs

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-unicode
// TODO: C.hb_buffer_serialize_unicode

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-deserialize-unicode
// TODO: C.hb_buffer_deserialize_unicode

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-format-from-string
// TODO: C.hb_buffer_serialize_format_from_string

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-format-to-string
// TODO: C.hb_buffer_serialize_format_to_string

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-list-formats
// TODO: C.hb_buffer_serialize_list_formats

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-segment-properties-equal
// TODO: C.hb_segment_properties_equal

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-segment-properties-hash
// TODO: C.hb_segment_properties_hash

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-segment-properties-overlay
// TODO: C.hb_segment_properties_overlay

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-diff
// TODO: C.hb_buffer_diff

// https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-message-func-t
// TODO: experiment what's hb_buffer_message_func_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-message-func
// TODO: C.hb_buffer_set_message_func
