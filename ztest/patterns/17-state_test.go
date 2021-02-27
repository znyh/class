package patterns

import (
	"testing"
)

func TestState(t *testing.T) {

	h := new(hero)

	h.state = new(concretestate)

	h.sethealth(5)
	h.showinfo()

	h.sethealth(75)
	h.showinfo()

	h.sethealth(95)
	h.showinfo()
}
