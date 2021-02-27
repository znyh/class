package patterns

import (
	"fmt"
)

type isoftware interface {
	running()
}

type cpusoft struct {
}

func (c *cpusoft) running() {
	fmt.Println("cpu is running")
}

type storagesoft struct {
}

func (c *storagesoft) running() {
	fmt.Println("storagesoft is running")
}

type bridge struct {
	soft isoftware
}

func (b *bridge) setsoftware(soft isoftware) {
	b.soft = soft
}

type phone struct {
	bridge
}

func (p *phone) setbridge(soft isoftware) {
	p.bridge.setsoftware(soft)
}

func (p *phone) work() {
	p.bridge.soft.running()
}
