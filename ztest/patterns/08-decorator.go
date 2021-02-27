package patterns

type idecorator interface {
	getnum() int
	getprice() int
}

type fruit struct {
	num   int
	price int
}

func (f *fruit) getnum() int {
	return f.num
}
func (f *fruit) getprice() int {
	return f.price
}

type appledecorator struct {
	idecorator
	addnum   int
	addprice int
}

func (a *appledecorator) getnum() int {
	return a.idecorator.getnum() + a.addnum
}
func (a *appledecorator) getprice() int {
	return a.idecorator.getprice() + a.addprice
}
