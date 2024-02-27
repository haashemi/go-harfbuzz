package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"

import (
	"unsafe"
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#HB-SEGMENT-PROPERTIES-DEFAULT:CAPS
const BufferReplacementCodepointDefault = 0xFFFD

// Buffer is the main structure holding the input text and its properties
// before shaping, and output glyphs and their information after shaping.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-t
type Buffer *C.hb_buffer_t

// GlyphInfo holds information about the glyphs and their relation to input text.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-glyph-info-t
type GlyphInfo struct {
	Codepoint uint32 // Either a Unicode code point (before shaping) or a glyph index (after shaping).
	mask      uint32 // [private]
	// The index of the character in the original text that corresponds to this
	// GlyphInfo, or whatever the client passes to BufferAdd. More than one
	// GlyphInfo can have the same cluster value, if they resulted from the same
	// character (e.g. one to many glyph substitution), and when more than one
	// character gets merged in the same glyph (e.g. many to one glyph substitution)
	// the hb_glyph_info_t will have the smallest cluster value of them. By default
	// some characters are merged into the same cluster (e.g. combining marks have
	// the same cluster as their bases) even if they are separate glyphs,
	// BufferSetClusterLevel allow selecting more fine-grained cluster handling.
	Cluster uint32
	var1    [4]byte // [private]
	var2    [4]byte // [private]
}

// GlyphFlags are flags for GlyphInfo.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-glyph-flags-t
type GlyphFlags C.hb_glyph_flags_t

const (
	GlyphFlagUnsafeToBreak       GlyphFlags = C.HB_GLYPH_FLAG_UNSAFE_TO_BREAK
	GlyphFlagUnsafeToConcat      GlyphFlags = C.HB_GLYPH_FLAG_UNSAFE_TO_CONCAT
	GlyphFlagSafeToInsertTatweel GlyphFlags = C.HB_GLYPH_FLAG_SAFE_TO_INSERT_TATWEEL
	GlyphFlagDefined             GlyphFlags = C.HB_GLYPH_FLAG_DEFINED
)

// GlyphPosition holds the positions of the glyph in both horizontal and vertical
// directions. All positions are relative to the current point.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-glyph-position-t
type GlyphPosition struct {
	XAdvance int32   // how much the line advances after drawing this glyph when setting text in horizontal direction.
	YAdvance int32   // how much the line advances after drawing this glyph when setting text in vertical direction.
	XOffset  int32   // how much the glyph moves on the X-axis before drawing it, this should not affect how much the line advances.
	YOffset  int32   // how much the glyph moves on the Y-axis before drawing it, this should not affect how much the line advances.
	var1     [4]byte // [private]
}

// ContentType is type of Buffer contents.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-content-type-t
type ContentType C.hb_buffer_content_type_t

const (
	ContentTypeInvalid ContentType = C.HB_BUFFER_CONTENT_TYPE_INVALID // Initial value for new buffer.
	ContentTypeUnicode ContentType = C.HB_BUFFER_CONTENT_TYPE_UNICODE // The buffer contains input characters (before shaping).
	ContentTypeGlyphs  ContentType = C.HB_BUFFER_CONTENT_TYPE_GLYPHS  // The buffer contains output glyphs (after shaping).
)

// BufferFlags are flags for Buffer contents.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-flags-t
type BufferFlags C.hb_buffer_flags_t

const (
	BufferFlagDefault                    BufferFlags = C.HB_BUFFER_FLAG_DEFAULT
	BufferFlagBot                        BufferFlags = C.HB_BUFFER_FLAG_BOT
	BufferFlagEot                        BufferFlags = C.HB_BUFFER_FLAG_EOT
	BufferFlagPreserveDefaultIgnorables  BufferFlags = C.HB_BUFFER_FLAG_PRESERVE_DEFAULT_IGNORABLES
	BufferFlagRemoveDefaultIgnorables    BufferFlags = C.HB_BUFFER_FLAG_REMOVE_DEFAULT_IGNORABLES
	BufferFlagDoNotInsertDottedCircle    BufferFlags = C.HB_BUFFER_FLAG_DO_NOT_INSERT_DOTTED_CIRCLE
	BufferFlagVerify                     BufferFlags = C.HB_BUFFER_FLAG_VERIFY
	BufferFlagProduceUnsafeToConcat      BufferFlags = C.HB_BUFFER_FLAG_PRODUCE_UNSAFE_TO_CONCAT
	BufferFlagProduceSafeToInsertTatweel BufferFlags = C.HB_BUFFER_FLAG_PRODUCE_SAFE_TO_INSERT_TATWEEL
	BufferFlagDEFINED                    BufferFlags = C.HB_BUFFER_FLAG_DEFINED
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-cluster-level-t
type ClusterLevel C.hb_buffer_cluster_level_t

const (
	ClusterLevelMonotoneGraphemes  ClusterLevel = C.HB_BUFFER_CLUSTER_LEVEL_MONOTONE_GRAPHEMES  // Return cluster values grouped by graphemes into monotone order.
	ClusterLevelMonotoneCharacters ClusterLevel = C.HB_BUFFER_CLUSTER_LEVEL_MONOTONE_CHARACTERS // Return cluster values grouped into monotone order.
	ClusterLevelCharacters         ClusterLevel = C.HB_BUFFER_CLUSTER_LEVEL_CHARACTERS          // Don't group cluster values.
	ClusterLevelDefault            ClusterLevel = C.HB_BUFFER_CLUSTER_LEVEL_DEFAULT             // Default cluster level, equal to ClusterLevelMonotoneGraphemes.
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-segment-properties-t
type SegmentProperties struct {
	Direction Direction      // the Direction of the buffer, see BufferSetDirection.
	Script    Script         // the Script of the buffer, see BufferSetScript.
	Language  Language       // the Language of the buffer, see BufferSetLanguage.
	reserved1 unsafe.Pointer // [private]
	reserved2 unsafe.Pointer // [private]
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-format-t
type SerializeFormat C.hb_buffer_serialize_format_t

const (
	SerializeFormatText    SerializeFormat = C.HB_BUFFER_SERIALIZE_FORMAT_TEXT    // a human-readable, plain text format.
	SerializeFormatJson    SerializeFormat = C.HB_BUFFER_SERIALIZE_FORMAT_JSON    // a machine-readable JSON format.
	SerializeFormatInvalid SerializeFormat = C.HB_BUFFER_SERIALIZE_FORMAT_INVALID // invalid format.
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-serialize-flags-t
type SerializeFlags C.hb_buffer_serialize_flags_t

const (
	SerializeFlagDefault      SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_DEFAULT        //serialize glyph names, clusters and positions.
	SerializeFlagNoClusters   SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_NO_CLUSTERS    //do not serialize glyph cluster.
	SerializeFlagNoPositions  SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_NO_POSITIONS   //do not serialize glyph position information.
	SerializeFlagNoGlyphNames SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_NO_GLYPH_NAMES //do no serialize glyph name.
	SerializeFlagGlyphExtents SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_GLYPH_EXTENTS  //serialize glyph extents.
	SerializeFlagGlyphFlags   SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_GLYPH_FLAGS    //serialize glyph flags.
	SerializeFlagNoAdvances   SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_NO_ADVANCES    //do not serialize glyph advances, glyph offsets will reflect absolute glyph positions.
	SerializeFlagDefined      SerializeFlags = C.HB_BUFFER_SERIALIZE_FLAG_DEFINED        //All currently defined flags.
)

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-diff-flags-t
type BufferDiffFlags C.hb_buffer_diff_flags_t

const (
	BufferDiffFlagEqual               BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_EQUAL                 // equal buffers.
	BufferDiffFlagContentTypeMismatch BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_CONTENT_TYPE_MISMATCH // buffers with different ContentType.
	BufferDiffFlagLengthMismatch      BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_LENGTH_MISMATCH       // buffers with differing length.
	BufferDiffFlagNotdefPresent       BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_NOTDEF_PRESENT        // .notdef glyph is present in the reference buffer.
	BufferDiffFlagDottedCirclePresent BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_DOTTED_CIRCLE_PRESENT // dotted circle glyph is present in the reference buffer.
	BufferDiffFlagCodepointPresent    BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_CODEPOINT_MISMATCH    // difference in GlyphInfo.Codepoint
	BufferDiffFlagClusterMismatch     BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_CLUSTER_MISMATCH      // difference in GlyphInfo.Cluster
	BufferDiffFlagGlyphFlagsMismatch  BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_GLYPH_FLAGS_MISMATCH  // difference in GlyphFlags.
	BufferDiffFlagPositionsMismatch   BufferDiffFlags = C.HB_BUFFER_DIFF_FLAG_POSITION_MISMATCH     // difference in GlyphPosition.
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

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-set-user-data
func BufferSetUserData(buffer Buffer, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_buffer_set_user_data(buffer, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-user-data
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
func BufferGetGlyphInfos(buffer Buffer) []GlyphInfo {
	var length uint32
	data := C.hb_buffer_get_glyph_infos(buffer, (*C.uint)(&length))

	res := make([]GlyphInfo, length)
	size := unsafe.Sizeof(C.hb_glyph_info_t{})

	for i := uint32(0); i < length; i++ {
		res[i] = *(*GlyphInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + size*uintptr(i)))
	}
	return res
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-glyph-info-get-glyph-flags
// TODO: C.hb_glyph_info_get_glyph_flags

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-buffer.html#hb-buffer-get-glyph-positions
func BufferGetGlyphPositions(buffer Buffer) []GlyphPosition {
	var length uint32
	data := C.hb_buffer_get_glyph_positions(buffer, (*C.uint)(&length))

	res := make([]GlyphPosition, length)
	size := unsafe.Sizeof(C.hb_glyph_position_t{})

	for i := uint32(0); i < length; i++ {
		res[i] = *(*GlyphPosition)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + size*uintptr(i)))
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
