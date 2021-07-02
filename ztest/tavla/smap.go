package calc

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/pkg/log"
)

const (
	_maxStoneNum  = 30
	_maxCampStone = 15
	_maxRow       = 25
)

const (
	ErrCodeNoExistStoneId              = 1001
	ErrCodeMoveByInBagStone            = 1002
	ErrCodeMoveByStoneNotInRow         = 1003
	ErrCodeMoveByStoneOutMaxRow        = 1004
	ErrCodeMoveByAllNoInHome           = 1005
	ErrCodeMoveByAtLeastTwoOtherStones = 1006
)

type Map struct {
	stones []Stone
	steps  []*Step
}

func NewMap() *Map {
	m := &Map{}
	for id := int32(0); id < _maxStoneNum; id++ {
		var s Stone
		s.init(id)
		m.stones = append(m.stones, s)
	}
	return m
}

func (m *Map) Stones() []Stone {
	return m.stones
}

func (m *Map) Steps() []*Step {
	return m.steps
}

func (m *Map) validStoneID(nId int32) bool {
	return nId >= 0 && nId < _maxStoneNum
}

func (m *Map) CanMove(moves []int32) (ok bool, code int32) {
	for i := 0; i < len(moves); i = i + 2 {
		if can, code := m.canMoveOne(moves[i], moves[i+1]); !can {
			return false, code
		}
	}
	return true, 0
}

func (m *Map) canMoveOne(moveID, x int32) (can bool, code int32) {

	if ok := m.validStoneID(moveID); !ok {
		code = ErrCodeNoExistStoneId
		return
	}

	c := m.stones[moveID]

	if c.inBag() {
		code = ErrCodeMoveByInBagStone
		return
	}

	if inBag, to := c.calcMoveTo(x); inBag {
		if !m.allInHome(c.emType) {
			code = ErrCodeMoveByStoneOutMaxRow
			return
		}

	} else {
		if others := m.calcOtherStoneAtRow(to, c.emType); len(others) >= 2 {
			code = ErrCodeMoveByAtLeastTwoOtherStones
			return
		}
	}

	return true, 0
}

func (m *Map) Move(moves []int32) (steps []*Step) {
	for i := 0; i < len(moves); i = i + 2 {
		step := m.moveOne(moves[i], moves[i+1])
		steps = append(steps, step)
	}
	return
}

func (m *Map) moveOne(moveID, x int32) (step *Step) {
	moverFrom, moverTo, killID := m.moveStone(moveID, x)
	killFrom, killTo := m.killStone(killID)
	step = m.saveStep(moveID, killID, x, moverFrom, moverTo, killFrom, killTo)
	return
}

func (m *Map) moveStone(moveID, x int32) (from, to int32, killID int32) {
	killID = -1
	c := m.stones[moveID]
	if inBag, to := c.calcMoveTo(x); !inBag {
		if others := m.calcOtherStoneAtRow(to, c.emType); len(others) == 1 {
			killID = others[0].nId
		}
	}

	from, to = m.stones[moveID].move(x)
	return
}

func (m *Map) killStone(killID int32) (from, to int32) {
	if m.validStoneID(killID) {
		from, to = m.stones[killID].eat()
		return
	}
	return -1, -1
}

func (m *Map) saveStep(moveID, killID, moveX, moveFrom, moveTo, killFrom, killTo int32) *Step {
	step := newStep(moveID, killID, moveX, moveFrom, moveTo, killFrom, killTo)
	m.steps = append(m.steps, step)
	return step
}

func (m *Map) allInHome(emType int32) bool {
	for _, v := range m.stones {
		if v.emType != emType {
			continue
		}
		if !v.inHome() {
			return false
		}
	}
	return true
}

func (m *Map) backOne() {
	s := m.steps[len(m.steps)-1]
	m.stones[s.moveId].nRow = s.moveFrom
	if m.validStoneID(s.killId) {
		m.stones[s.killId].nRow = s.killFrom
	}

	m.steps = m.steps[:len(m.steps)-1]
}

func (m *Map) calcOtherStoneAtRow(nRow, emType int32) (others []Stone) {
	for _, v := range m.stones {
		if v.emType == emType {
			continue
		}
		if v.nRow == nRow {
			others = append(others, v)
		}
	}
	return
}

func (m *Map) Over() (over bool, emType int32) {
	cm := map[int32]int{}
	for _, v := range m.stones {
		if v.nRow == _maxRow && v.emType == EmWHITE {
			cm[EmWHITE]++
		}
		if v.nRow == 0 && v.emType == EmBLACK {
			cm[EmBLACK]++
		}
	}
	for k, v := range cm {
		if v >= _maxCampStone {
			return true, k
		}
	}
	return
}

//可移动路径的全排列: 棋子类型为emType，可移动色子点数为dices的 所有可移动的路径组合
func (m *Map) Permute(emType int32, dices []int32) [][]int32 {
	if len(dices) == 0 {
		return nil
	}
	var res [][]int32
	var cache []int32
	var visited = make([]bool, len(dices))
	m.tryFind(&res, cache, emType, dices, visited)

	res1 := [][]int32(nil)
	res2 := [][]int32(nil)
	res3 := [][]int32(nil)

	for _, v := range res {
		//一步的路径
		if len(v) == 2 {
			if ok, _ := m.canMoveOne(v[0], v[1]); ok {
				res1 = append(res1, v)
			}
		}

		//二步的路径
		if len(v) == 4 {
			if ok1, _ := m.canMoveOne(v[0], v[1]); ok1 {
				m.moveOne(v[0], v[1])
				if ok2, _ := m.canMoveOne(v[2], v[3]); ok2 {
					m.moveOne(v[2], v[3])
					m.backOne()
					res2 = append(res2, v) // ok && ok
				}
				m.backOne()
			}
		}

	}

	if dices[0] == dices[1] {
		_res3:=
	}

	log.Info("Permute, emType:%+v dices:%+v cnt:%d cnt1:%d cnt2:%d", emType, dices, len(res), len(res1), len(res2))

	if len(res2) == 0 {
		return res1
	}

	return res2
}

func (m *Map) tryFind(res *[][]int32, cache []int32, emType int32, dices []int32, visited []bool) {
	// 成功找到一组
	if len(cache) == len(dices) || len(cache)/2 == len(dices) {
		var c = make([]int32, len(cache))
		copy(c, cache)
		*res = append(*res, c)
		if len(cache)/2 == len(dices) {
			return
		}
	}

	// 回溯
	for j := 0; j < len(dices); j++ {

		for i := int32(0); i < int32(len(m.stones)); i++ {

			//剪枝
			if m.stones[i].emType != emType {
				continue
			}

			if !visited[j] {
				visited[j] = true

				////能够移动?
				//ok, _ := m.canMoveOne(i, dices[j])
				//if !ok {
				//	continue
				//}
				////尝试移动一次
				//m.moveOne(i, dices[j])

				cache = append(cache, i, dices[j])
				m.tryFind(res, cache, emType, dices, visited)
				cache = cache[:len(cache)-2]

				////回退之前移动的位置
				//m.backOne()

				visited[j] = false
			}
		}
	}
}

type TagRow struct {
	Count  int32
	EmType int32
}

func (m *Map) Show() string {
	cm := m.ToCountMap()
	b := strings.Builder{}

	for i := int32(0); i <= _maxRow; i++ {
		if val, ok := cm[i]; ok && val.Count > 0 {
			str := "W"
			if cm[i].EmType == EmBLACK {
				str = "B"
			}
			b.WriteString(fmt.Sprintf("|%s%d|", str, cm[i].Count))
		} else {
			b.WriteString(fmt.Sprintf("[%d]", i))
		}
	}
	return b.String()
}

func (m *Map) ToCountMap() map[int32]*TagRow {
	cm := map[int32]*TagRow{} // key:nRow val:stone
	for i := int32(0); i <= _maxRow; i++ {
		cm[i] = &TagRow{
			Count:  0,
			EmType: -1,
		}
	}
	for _, v := range m.stones {
		if c, ok := cm[v.nRow]; ok {
			cm[v.nRow].Count++
			cm[v.nRow].EmType = c.EmType
		}
	}

	return cm
}
