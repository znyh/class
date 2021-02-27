package patterns

type itemplate interface {
	getname() string
	setname(string)
}

type component struct {
	name string
	age  int
}

func (c *component) getname() string {
	return c.name
}

func (c *component) setname(name string) {
	c.name = name
}

type template struct {
	*component
}

func (c *template) setname(name string) {
	c.component.name = name
}
