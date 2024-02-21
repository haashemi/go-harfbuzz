package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-t
type Font *C.hb_font_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-add-glyph-origin-for-direction
func FontAddGlyphOriginForDirection(font Font, glyph uint32, direction Direction, x, y *int) {
	C.hb_font_add_glyph_origin_for_direction(font, C.uint(glyph), C.hb_direction_t(direction), (*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-create
func FontCreate(face Face) Font {
	return C.hb_font_create(face)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-create-sub-font
func FontCreateSubFont(parent Font) Font {
	return C.hb_font_create_sub_font(parent)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-get-empty
func FontGetEmpty() Font {
	return C.hb_font_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-reference
func FontReference(font Font) Font {
	return C.hb_font_reference(font)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-destroy
func FontDestroy(font Font) {
	C.hb_font_destroy(font)
}

// TODO: hb_font_set_user_data
// TODO: hb_font_get_user_data
// TODO: hb_font_make_immutable
// TODO: hb_font_is_immutable
// TODO: hb_font_set_face
// TODO: hb_font_get_face
// TODO: hb_font_get_glyph
// TODO: hb_font_get_glyph_advance_for_direction
// TODO: hb_font_get_glyph_advances_for_direction
// TODO: hb_font_get_glyph_contour_point
// TODO: hb_font_get_glyph_contour_point_for_origin
// TODO: hb_font_get_glyph_extents
// TODO: hb_font_get_glyph_extents_for_origin
// TODO: hb_font_get_glyph_from_name
// TODO: hb_font_get_glyph_h_advance
// TODO: hb_font_get_glyph_v_advance
// TODO: hb_font_get_glyph_h_advances
// TODO: hb_font_get_glyph_v_advances
// TODO: hb_font_get_glyph_h_kerning
// TODO: hb_font_get_glyph_kerning_for_direction
// TODO: hb_font_get_glyph_h_origin
// TODO: hb_font_get_glyph_v_origin
// TODO: hb_font_get_glyph_origin_for_direction
// TODO: hb_font_get_glyph_name
// TODO: hb_font_draw_glyph
// TODO: hb_font_paint_glyph
// TODO: hb_font_get_nominal_glyph
// TODO: hb_font_get_nominal_glyphs
// TODO: hb_font_get_variation_glyph
// TODO: hb_font_set_parent
// TODO: hb_font_get_parent
// TODO: hb_font_set_ppem
// TODO: hb_font_get_ppem
// TODO: hb_font_set_ptem
// TODO: hb_font_get_ptem
// TODO: hb_font_set_scale
// TODO: hb_font_get_scale
// TODO: hb_font_get_synthetic_bold
// TODO: hb_font_set_synthetic_bold
// TODO: hb_font_set_synthetic_slant
// TODO: hb_font_get_synthetic_slant
// TODO: hb_font_set_variations
// TODO: hb_font_set_variation
// TODO: hb_font_set_var_named_instance
// TODO: hb_font_get_var_named_instance
// TODO: hb_font_set_var_coords_design
// TODO: hb_font_get_var_coords_design
// TODO: hb_font_set_var_coords_normalized
// TODO: hb_font_get_var_coords_normalized
// TODO: hb_font_glyph_from_string
// TODO: hb_font_glyph_to_string
// TODO: hb_font_get_serial
// TODO: hb_font_changed
// TODO: hb_font_set_funcs
// TODO: hb_font_set_funcs_data
// TODO: hb_font_subtract_glyph_origin_for_direction
// TODO: hb_font_funcs_create
// TODO: hb_font_funcs_get_empty
// TODO: hb_font_funcs_reference
// TODO: hb_font_funcs_destroy
// TODO: hb_font_funcs_set_user_data
// TODO: hb_font_funcs_get_user_data
// TODO: hb_font_funcs_make_immutable
// TODO: hb_font_funcs_is_immutable
// TODO: (*hb_font_get_glyph_contour_point_func_t)
// TODO: hb_font_funcs_set_glyph_contour_point_func
// TODO: (*hb_font_get_glyph_extents_func_t)
// TODO: hb_font_funcs_set_glyph_extents_func
// TODO: (*hb_font_get_glyph_from_name_func_t)
// TODO: hb_font_funcs_set_glyph_from_name_func
// TODO: (*hb_font_get_glyph_advance_func_t)
// TODO: hb_font_funcs_set_glyph_h_advance_func
// TODO: hb_font_funcs_set_glyph_v_advance_func
// TODO: (*hb_font_get_glyph_advances_func_t)
// TODO: hb_font_funcs_set_glyph_h_advances_func
// TODO: hb_font_funcs_set_glyph_v_advances_func
// TODO: (*hb_font_get_glyph_kerning_func_t)
// TODO: hb_font_funcs_set_glyph_h_kerning_func
// TODO: (*hb_font_get_glyph_origin_func_t)
// TODO: hb_font_funcs_set_glyph_h_origin_func
// TODO: hb_font_funcs_set_glyph_v_origin_func
// TODO: (*hb_font_get_glyph_name_func_t)
// TODO: hb_font_funcs_set_glyph_name_func
// TODO: (*hb_font_draw_glyph_func_t)
// TODO: hb_font_funcs_set_draw_glyph_func
// TODO: (*hb_font_paint_glyph_func_t)
// TODO: hb_font_funcs_set_paint_glyph_func
// TODO: (*hb_font_get_nominal_glyph_func_t)
// TODO: hb_font_funcs_set_nominal_glyph_func
// TODO: (*hb_font_get_nominal_glyphs_func_t)
// TODO: hb_font_funcs_set_nominal_glyphs_func
// TODO: (*hb_font_get_variation_glyph_func_t)
// TODO: hb_font_funcs_set_variation_glyph_func

// A callback function for FaceCreateForTables.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-reference-table-func-t
type ReferenceTableFunc C.hb_reference_table_func_t

// TODO: (*hb_font_get_font_extents_func_t)
// TODO: hb_font_funcs_set_font_h_extents_func
// TODO: hb_font_funcs_set_font_v_extents_func
// TODO: hb_font_get_h_extents
// TODO: hb_font_get_v_extents
// TODO: hb_font_get_extents_for_direction
