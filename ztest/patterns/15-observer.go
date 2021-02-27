package patterns

import (
	"fmt"
)

type iobserver interface {
	receive(event)
}

type ishader interface {
	register(...iobserver)
	remove(...iobserver)
	notify(event)
}

type event struct {
	level int
	info  string
}

type observerA struct {
	name string
}

func (o *observerA) receive(e event) {
	fmt.Printf("observerA: %s接收到事件通知:[level:%d,info:%s]\n", o.name, e.level, e.info)
}

type observerB struct {
	name string
}

func (o *observerB) receive(e event) {
	fmt.Printf("observerB: %s接收到事件通知:[level:%d,info:%s]\n", o.name, e.level, e.info)
}

type shader struct {
	list []iobserver //*list.List ???
}

func (s *shader) register(obs ...iobserver) {
	if s.list == nil {
		s.list = make([]iobserver, 0)
	}

	for _, ob := range obs {
		if ob != nil {
			s.list = append(s.list, ob)
		}
	}
}
func (s *shader) remove(obs ...iobserver) {
	if s.list == nil || len(s.list) == 0 {
		return
	}

	for _, ob := range obs {
		for index, v := range s.list {
			if v == ob {
				s.list = append(s.list[:index], s.list[index+1:]...)
			}
		}
	}
}
func (s *shader) notify(e event) {
	for _, ob := range s.list {
		ob.receive(e)
	}
}
