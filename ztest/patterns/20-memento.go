package patterns

type origin struct {
	state int
}

type memento struct {
	o origin
}

type caretaker struct {
}

func (c *caretaker) creatememento(o origin) memento {
	return memento{o: o}
}

func (c *caretaker) recoverorigin(m memento) origin {
	return m.o
}
