package patterns

type prototype struct {
	desc string
}

func (p *prototype) clone() *prototype {
	tmp := *p
	return &tmp
}
