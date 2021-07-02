package calc

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

func TestRandTwoDices(t *testing.T) {

	const (
		_cnt = 10
	)
	for i := 0; i < _cnt; i++ {
		//不能相等
		ds, max := RandTwoDices(false)
		log.Info("cm:%+v,max:%d", ds, max)
	}
}

//3)以持有黑子为例，15枚棋子分布在：24点上2个，13点上5个，8点上3个，6点上5个。
//4)以持白子为例，15枚棋子分布在：1点上2个，12点上5个，17点上3个，19点上5个。
func TestNewMap(t *testing.T) {

	var (
		id    int32 = 15
		x     int32 = 3
		moves       = []int32{id, x}
		steps       = []*Step(nil)
	)

	m := NewMap()

	log.Info("Map:%+v", m.Show())

	ok, code := m.CanMove(moves)
	if ok {
		steps = m.Move(moves)
	}

	log.Info("code:%d x:%+v steps:%+v", code, x, DescSteps(steps))

	log.Info("Map:%+v", m.Show())

	m.backOne()

	log.Info("Map:%+v", m.Show())

}

func TestPermute(t *testing.T) {
	var (
		m      = NewMap()
		emType = EmWHITE
		start  = time.Now()
		dices  = []int32{6, 6} //2个色子
	)
	log.Info("Map:%+v", m.Show())
	ret := m.Permute(int32(emType), dices)
	str := "\n"

	for _, v := range ret {
		if len(v) == 2 {
			str += fmt.Sprintf("{ID:%d->%d}\n", v[0], v[1])
		} else if len(v) == 4 {
			str += fmt.Sprintf("{ID:%d->%d, ID:%d->%d}\n", v[0], v[1], v[2], v[3])
		} else {
			log.Error("====> bad step, %+v", v)
		}
	}
	log.Info("Map:%+v", m.Show())
	log.Info("use:%v/s cnt:%+v dices:%+v %+v", time.Since(start).Seconds(), len(ret), dices, str)

}
