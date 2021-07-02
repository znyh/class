package calc

import (
	"fmt"

	"class/ztest/base"
)

//随机两个色子 色子点数范围[1,6]
func RandTwoDices(canEqual bool) (dices []int32, maxIdx int32) {
	d1 := base.RandRange(1, 7)
	d2 := base.RandRange(1, 7)
	for !canEqual && (d1 == d2) {
		d1 = base.RandRange(1, 7)
		d2 = base.RandRange(1, 7)
	}
	dices = append(dices, int32(d1), int32(d2))
	if d2 > d1 {
		maxIdx = 1
	}
	return
}

func DescSteps(steps []*Step) (s string) {
	s = "【"
	for _, v := range steps {
		s += fmt.Sprintf("%+v, ", v.Desc())
	}
	s += "】"
	return
}
