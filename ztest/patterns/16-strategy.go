package patterns

type ioperator interface {
	operate(int, int) int
}

type addition struct {
}

func (add *addition) operate(a, b int) int {
	return a + b
}

type multiplication struct {
}

func (multi *multiplication) operate(a, b int) int {
	return a * b
}

type operation struct {
	ioperator
}

func (oper *operation) operate(o ioperator, a, b int) int {
	oper.ioperator = o
	return oper.ioperator.operate(a, b)
}
