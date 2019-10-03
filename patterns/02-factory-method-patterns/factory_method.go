package factory

import "errors"

/*
	工厂方法模式:使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化

	设计思想：
		*类型常量
		*接口factory
		*生成函数
		*实现接口方法的struct
*/

type Payment interface {
	Pay(money float32) error
}

type Kind int

const (
	Cash Kind = 1 << iota
	Credit
)

//实现两个struct,继承接口Payment
type CashPay struct {
	Balance float32
}

type CreditPay struct {
	Balance float32
}

func (cash *CashPay) Pay(money float32) error {
	if cash.Balance < 0 || cash.Balance < money {
		return errors.New("balance not enough")
	}
	cash.Balance -= money
	return nil
}

func (credit *CreditPay) Pay(money float32) error {
	if credit.Balance < 0 || credit.Balance < money {
		return errors.New("balance not enough")
	}
	credit.Balance -= money
	return nil
}

func GeneratePayment(k Kind, balance float32) (Payment, error) {
	switch k {
	case Cash:
		cash := new(CashPay)
		cash.Balance = balance
		return cash, nil
	case Credit:
		return &CreditPay{balance}, nil
	default:
		return nil, errors.New("Payment do not support this ")
	}
}
