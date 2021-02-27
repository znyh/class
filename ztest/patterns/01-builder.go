package patterns

type ibuilder interface {
	setwheel() ibuilder
	setseat() ibuilder
	getvehicle() vehicle
}

type vehicle struct {
	wheel int

	seat int
}

type car struct {
	vehicle
}

func (c *car) setwheel() ibuilder {
	c.wheel = 4
	return c
}

func (c *car) setseat() ibuilder {
	c.seat = 4
	return c
}

func (c *car) getvehicle() vehicle {
	return c.vehicle
}

type director struct {
	ibuilder //这样，可以重写特定的方法而不必定义所有其他方法
}

func (d *director) construct(b ibuilder) {
	b.setwheel().setseat()
	d.ibuilder = b
}

func (d *director) setseat() ibuilder {
	d.ibuilder.setseat()
	return d
}
