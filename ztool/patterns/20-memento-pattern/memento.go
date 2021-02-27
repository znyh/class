package memento

/*
	备忘录模式: 建立原始对象副本，用于存储恢复原始对象数据。
*/

type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

type Memento struct {
	Originator
}

func (m *Memento) GetState() string {
	return m.Originator.state
}

func (m *Memento) SetState(originator Originator) {
	m.Originator = originator
}

type Caretaker struct {
	//memento Memento
}

func (c *Caretaker) CreateMemento(originator Originator) Memento {
	return Memento{originator}
}

func (c *Caretaker) RecoverOriginator(memento Memento) Originator {
	return memento.Originator
}
