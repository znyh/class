package patterns

import (
	"fmt"
)

type iproduce interface {
	describe()
}
type ifactory interface {
	produce() iproduce
}

type concreateproduce struct {
	name string
}

func (p *concreateproduce) describe() {
	fmt.Println("this is a produce,name:", p.name)
}

type concreatefactory struct {
}

func (f *concreatefactory) produce() iproduce {
	return &concreateproduce{name: "kd"}
}
