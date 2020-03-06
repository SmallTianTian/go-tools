package object

import (
	"strconv"
	"strings"
)

func StringIsBlank(content string) bool {
	return content == "" || strings.TrimSpace(content) == ""
}

func MustUint64(value string) uint64 {
	if v, err := strconv.ParseUint(value, 10, 64); err != nil {
		panic(err)
	} else {
		return v
	}
}

func MustInt64(value string) int64 {
	if v, err := strconv.ParseInt(value, 10, 64); err != nil {
		panic(err)
	} else {
		return v
	}
}
