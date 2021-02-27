package mediator

import "fmt"

/*
	中介者模式：将对象之间的通信关联关系封装到一个中介类中单独处理，
                从而使其耦合松散，可以独立地改变它们之间的交互
*/

type IDepartment interface {
	SendMess(message string)
	GetMess(message string)
}

type Mediator struct {
	Market
	Technical
}

//消息转发
func (m *Mediator) ForwardMessage(department IDepartment, message string) {
	switch department.(type) {
	case *Technical:
		m.Market.GetMess(message)
	case *Market:
		m.Technical.GetMess(message)
	default:
		fmt.Println("部门不在中介者中")
	}
}

type Technical struct {
	mediator *Mediator
}

func (t *Technical) SendMess(message string) {
	t.mediator.ForwardMessage(t, message)
}

func (t *Technical) GetMess(message string) {
	fmt.Printf("技术部收到消息: %s\n", message)
}

type Market struct {
	mediator *Mediator
}

func (m *Market) SendMess(message string) {
	m.mediator.ForwardMessage(m, message)
}

func (m *Market) GetMess(message string) {
	fmt.Printf("市场部部收到消息: %s\n", message)
}
