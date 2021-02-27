package patterns

type iflyweight interface {
	setradius(int)
	setcolor(int)
}

type circle struct {
	radius int
	color  int
}

func (c *circle) setradius(r int) {
	c.radius = r
}
func (c *circle) setcolor(color int) {
	c.color = color
}

type circleFactory struct {
	list map[int]iflyweight
}

func (cf *circleFactory) getcircle(color int) iflyweight {
	if cf.list == nil {
		cf.list = make(map[int]iflyweight)
	}

	if v, ok := cf.list[color]; ok {
		return v
	}

	v := new(circle)
	v.setcolor(color)
	cf.list[color] = v

	return v
}
