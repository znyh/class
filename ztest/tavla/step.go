package calc

import "fmt"

type Step struct {
	moveId   int32
	moveX    int32
	moveFrom int32
	moveTo   int32
	killId   int32
	killFrom int32
	killTo   int32
}

func newStep(moveId, killId, moveX, moveFrom, moveTo, killFrom, killTo int32) *Step {
	var step = &Step{
		moveId:   moveId,
		moveX:    moveX,
		moveFrom: moveFrom,
		moveTo:   moveTo,
		killId:   killId,
		killFrom: killFrom,
		killTo:   killTo,
	}
	return step
}

func (s *Step) Desc() string {
	return fmt.Sprintf("%+v", s)
}
func (s *Step) GetMoveId() int32 {
	return s.moveId
}
func (s *Step) GetKillId() int32 {
	return s.killId
}
func (s *Step) GetMoveX() int32 {
	return s.moveX
}
func (s *Step) GetMoveFrom() int32 {
	return s.moveFrom
}
func (s *Step) GetMoveTo() int32 {
	return s.moveTo
}
func (s *Step) GetKillFrom() int32 {
	return s.killFrom
}
func (s *Step) GetKillTo() int32 {
	return s.killTo
}
