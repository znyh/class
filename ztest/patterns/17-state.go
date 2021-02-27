package patterns

import (
	"fmt"
)

type istate interface {
	showinfo()
	setinfo(string)
}

type concretestate struct {
	info string
}

func (c *concretestate) showinfo() {
	fmt.Println("info:", c.info)
}
func (c *concretestate) setinfo(info string) {
	c.info = info
}

type hero struct {
	health int
	state  istate
}

func (h *hero) showinfo() {
	h.state.showinfo()
}

func (h *hero) sethealth(health int) {
	h.health = health
	h.changestate()
}

func (h *hero) changestate() {
	if h.health <= 10 {
		h.state.setinfo("restricted_state")
	} else if h.health <= 80 {
		h.state.setinfo("normal_state")
	} else if h.health <= 100 {
		h.state.setinfo("close_state")
	}
}
