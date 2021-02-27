package patterns

import (
	"testing"
)

func TestObserver(t *testing.T) {

	s := new(shader)
	oba := &observerA{name: "a"}
	obb := &observerB{name: "b"}

	s.register(oba, obb)
	s.notify(event{level: 1, info: "notify.."})

	s.remove(oba)
	s.notify(event{level: 2, info: "a 已经离开了"})
}
