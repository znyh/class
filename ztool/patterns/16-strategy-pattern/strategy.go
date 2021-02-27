package strategy

/*
	策略模式: 定义多个不同的实现类，这些类实现公共接口，通过调用接口调用不同实例得到不同结果
*/

type Operator interface {
	Apply(int, int) int
}

type Addition struct{}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

type Multiplication struct{}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}

//包装器
type Operation struct {
	operator Operator
}

func (op *Operation) Operate(left, right int) int {
	return op.operator.Apply(left, right)
}

func CreateOpration(operator Operator) Operation {
	return Operation{operator}
}
