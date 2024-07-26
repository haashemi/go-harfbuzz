package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import "unsafe"

// ShapePlan holds a shaping plan.
//
// Shape plans contain information about how HarfBuzz will shape a particular
// text segment, based on the segment's properties and the capabilities in the
// font face in use.
//
// Shape plans can be queried about how shaping will perform, given a set of
// specific input parameters (script, language, direction, features, etc.).
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-t
type ShapePlan *C.hb_shape_plan_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-create
func ShapePlanCreate(face Face, props *SegmentProperties, userFeatures []Feature, shaperList []string) ShapePlan {
	shapers := cStringArray(shaperList)
	defer freeStringArray(shapers)

	return C.hb_shape_plan_create(face, (*C.hb_segment_properties_t)(unsafe.Pointer(props)), cFeatures(userFeatures), C.uint(len(userFeatures)), &shapers[0])
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-create-cached
func ShapePlanCreateCached(face Face, props *SegmentProperties, userFeatures []Feature, shaperList []string) ShapePlan {
	shapers := cStringArray(shaperList)
	defer freeStringArray(shapers)

	return C.hb_shape_plan_create_cached(face, (*C.hb_segment_properties_t)(unsafe.Pointer(props)), cFeatures(userFeatures), C.uint(len(userFeatures)), &shapers[0])
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-create2
func ShapePlanCreate2(face Face, props *SegmentProperties, userFeatures []Feature, coords []int32, shaperList []string) ShapePlan {
	var cCoords *C.int
	if len(coords) > 0 {
		cCoords = (*C.int)(unsafe.Pointer(&coords[0]))
	}

	shapers := cStringArray(shaperList)
	defer freeStringArray(shapers)

	return C.hb_shape_plan_create2(face, (*C.hb_segment_properties_t)(unsafe.Pointer(props)), cFeatures(userFeatures), C.uint(len(userFeatures)), cCoords, C.uint(len(coords)), &shapers[0])
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-create-cached2
func ShapePlanCreateCached2(face Face, props *SegmentProperties, userFeatures []Feature, coords []int32, shaperList []string) ShapePlan {
	var cCoords *C.int
	if len(coords) > 0 {
		cCoords = (*C.int)(unsafe.Pointer(&coords[0]))
	}

	shapers := cStringArray(shaperList)
	defer freeStringArray(shapers)

	return C.hb_shape_plan_create_cached2(face, (*C.hb_segment_properties_t)(unsafe.Pointer(props)), cFeatures(userFeatures), C.uint(len(userFeatures)), cCoords, C.uint(len(coords)), &shapers[0])
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-get-empty
func ShapePlanGetEmpty() ShapePlan {
	return C.hb_shape_plan_get_empty()
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-reference
func ShapePlanReference(shapePlan ShapePlan) ShapePlan {
	return C.hb_shape_plan_reference(shapePlan)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-destroy
func ShapePlanDestroy(shapePlan ShapePlan) {
	C.hb_shape_plan_destroy(shapePlan)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-set-user-data
func ShapePlanSetUserData(shapePlan ShapePlan, key *UserDataKey, data unsafe.Pointer, destroy DestroyFunc, replace bool) bool {
	return C.hb_shape_plan_set_user_data(shapePlan, (*C.hb_user_data_key_t)(key), data, destroy, cBool(replace)) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-get-user-data
func ShapePlanGetUserData(shapePlan ShapePlan, key *UserDataKey) unsafe.Pointer {
	return C.hb_shape_plan_get_user_data(shapePlan, (*C.hb_user_data_key_t)(key))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-execute
func ShapePlanExecute(shapePlan ShapePlan, font Font, buffer Buffer, features []Feature) bool {
	return C.hb_shape_plan_execute(shapePlan, font, buffer, cFeatures(features), C.uint(len(features))) == 1
}

// ShapePlanGetShaper returns the shaper from a given shaping plan.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-shape-plan.html#hb-shape-plan-get-shaper
func ShapePlanGetShaper(shapePlan ShapePlan) string {
	res := C.hb_shape_plan_get_shaper(shapePlan)
	return C.GoString(res)
}
