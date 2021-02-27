package patterns

import (
	"container/list"
	"fmt"
)

type ivisitor interface {
	visit()
	exit()
}

type visitor struct {
	name string
}

func (v *visitor) visit() {
	fmt.Printf("访问者%s正在访问\n", v.name)
}
func (v *visitor) exit() {
	fmt.Printf("访问者%s离开访问\n", v.name)
}

type door struct {
	list *list.List
}

func (d *door) accept(vs ...ivisitor) {
	if d.list == nil {
		d.list = list.New()
	}

	for _, v := range vs {
		d.list.PushBack(v)
		v.visit()
	}
}

func (d *door) remove(vs ...ivisitor) {
	if d.list.Len() <= 0 {
		return
	}

	for _, v := range vs {
		for e := d.list.Front(); e != nil; e = e.Next() {
			if v == e.Value {
				elem := d.list.Remove(e)
				elem.(*visitor).exit()
			}
		}
	}
}

func (d *door) close() {
	d.list.Init()
}
