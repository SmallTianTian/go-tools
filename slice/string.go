package slice

import "github.com/SmallTianTian/go-tools/cst"

func StringInSlice(slice []string, key string) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func DistinctStringSlice(origin []string) []string {
	if len(origin) == 0 {
		return origin
	}

	set := make(map[string]struct{})
	result := make([]string, 0, len(origin))

	for _, item := range origin {
		if _, in := set[item]; !in {
			result = append(result, item)
			set[item] = cst.EmptyStruct
		}
	}

	return result
}
