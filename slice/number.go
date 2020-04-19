package slice

import (
	"reflect"

	"github.com/SmallTianTian/go-tools/cst"
)

// --- in
func UintPtrInSlice(slice []uintptr, key uintptr) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Uint8InSlice(slice []uint8, key uint8) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Uint16InSlice(slice []uint16, key uint16) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Uint32InSlice(slice []uint32, key uint32) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func UintInSlice(slice []uint, key uint) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Uint64InSlice(slice []uint64, key uint64) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Int8InSlice(slice []int8, key int8) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Int16InSlice(slice []int16, key int16) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Int32InSlice(slice []int32, key int32) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func IntInSlice(slice []int, key int) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Int64InSlice(slice []int64, key int64) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Float32InSlice(slice []float32, key float32) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func Float64InSlice(slice []float64, key float64) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

// 是否在列表中。
// key 的类型，必须与 slice 中元素的类型一样，采用系统 == 方式，不会做深层次比较。
// @Param slice: 列表
// @Param key  : 将比较的值
func InSlice(slice, key interface{}) bool {
	if reflect.TypeOf(slice).Kind() != reflect.Slice && reflect.TypeOf(slice).Kind() != reflect.Array {
		return false
	}

	sv := reflect.ValueOf(slice)
	len := sv.Len()
	for i := 0; i < len; i++ {
		if sv.Index(i).Interface() == key {
			return true
		}
	}
	return false
}

// --- distinct
func DistinctUintptrSlice(origin []uintptr) []uintptr {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uintptr]struct{})
	result := make([]uintptr, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctUint8Slice(origin []uint8) []uint8 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uint8]struct{})
	result := make([]uint8, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctUint16Slice(origin []uint16) []uint16 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uint16]struct{})
	result := make([]uint16, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctUint32Slice(origin []uint32) []uint32 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uint32]struct{})
	result := make([]uint32, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctUintSlice(origin []uint) []uint {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uint]struct{})
	result := make([]uint, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctUint64Slice(origin []uint64) []uint64 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[uint64]struct{})
	result := make([]uint64, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctInt8Slice(origin []int8) []int8 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[int8]struct{})
	result := make([]int8, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctInt16Slice(origin []int16) []int16 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[int16]struct{})
	result := make([]int16, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctInt32Slice(origin []int32) []int32 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[int32]struct{})
	result := make([]int32, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctIntSlice(origin []int) []int {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[int]struct{})
	result := make([]int, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctInt64Slice(origin []int64) []int64 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[int64]struct{})
	result := make([]int64, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctFloat32Slice(origin []float32) []float32 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[float32]struct{})
	result := make([]float32, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

func DistinctFloat64Slice(origin []float64) []float64 {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[float64]struct{})
	result := make([]float64, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}

// 使用示例：
// a := {1, 2, 3, 1}
// c := DistinctSlice(a)
// fmt.Println(c) // {1, 2, 3}
func DistinctSlice(origin interface{}) interface{} {
	if reflect.TypeOf(origin).Kind() != reflect.Slice && reflect.TypeOf(origin).Kind() != reflect.Array {
		return origin
	}

	ov := reflect.ValueOf(origin)
	if ov.Len() == 0 {
		return origin
	}

	t := ov.Type()
	fresh := reflect.MakeSlice(t, 0, ov.Len())
	for i := 0; i < ov.Len(); i++ {
		k := ov.Index(i)
		in := false
		for j := 0; j < fresh.Len(); j++ {
			if k.Interface() == fresh.Index(j).Interface() {
				in = true
				break
			}
		}

		if !in {
			fresh = reflect.Append(fresh, k)
		}
	}

	return fresh.Interface()
}
