package util

import (
	"fmt"
	"strconv"
	"strings"
)

//ToString .
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//字符串表示的16进制转化为整形
func HexStr2Int32(src string) (dst []int32) {
	if len(src) <= 0 {
		return
	}

	tmp := strings.TrimSpace(src)
	s := strings.Split(tmp, ",")

	for _, v := range s {
		vv := strings.TrimSpace(v)
		if val, err := strconv.ParseInt(vv, 0, 0); err == nil {
			dst = append(dst, int32(val))
		} else {
			return []int32{}
		}
	}

	return dst
}
