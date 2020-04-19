package slice

import (
	"reflect"
)

// For example:
// c := []int{1, 2, 3, 4}
// RemoveSlice(&c, 1)
// fmt.Println(c) // {1, 3, 4}
func RemoveSlice(slice interface{}, i int) {
	st := reflect.TypeOf(slice)
	if st.Kind() != reflect.Ptr {
		return
	}

	sv := reflect.ValueOf(slice)
	sv = sv.Elem()
	if sv.Type().Kind() != reflect.Slice {
		return
	}

	l := sv.Len()
	switch true {
	case l == 0 || i >= l:
	case i == 0:
		sv.Set(sv.Slice(1, l))
	case i > 0 && i < l-1:
		sv.Set(reflect.AppendSlice(sv.Slice(0, i), sv.Slice(i+1, l)))
	case i == l-1:
		sv.Set(sv.Slice(0, i))
	}
}
