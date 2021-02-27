package patterns

import (
	"fmt"
)

type runaction struct {
}

func (o *runaction) doaction(action string) {
	fmt.Println("objectaction do action:", action)
}

type readaction struct {
}

func (o *readaction) doaction(action string) {
	fmt.Println("readaction do action:", action)
}

type actionproxy struct {
	run  runaction
	read readaction
}

func (p *actionproxy) doaction(action string) {
	if action == "run" {
		p.run.doaction(action)
	} else if action == "read" {
		p.read.doaction(action)
	}
}
