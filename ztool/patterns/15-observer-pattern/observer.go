package observer

import "fmt"

/*
	观察者模式定义对象间的一种一对多的依赖关系,以便当一个对象的状态发生改变时,所有依赖于它的对象都得到通知并自动刷新

	设计思想
		1. Event struct
		2. Observer interface
			OnNotify(Event) 处理事件
		3. 被观察者Notifier interface 实现以下三个方法
			实现 Register ObServer
			取消 DeRegister Observer
			通知 Notify(Event)
*/

type Observer interface {
	Receive(event Event)
}

type Notifier interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event Event)
}

type Event struct {
	Info string
}

func NewEvent() Event {
	return Event{Info: "价格变动通知"}
}

type InvestorObserver struct {
	Name string
}

func NewInvestorObserver(name string) *InvestorObserver {
	return &InvestorObserver{Name: name}
}

func (invester *InvestorObserver) Receive(event Event) {
	fmt.Printf("%s 收到事件通知 %s\n", invester.Name, event.Info)
}

type ShareNotifier struct {
	Price  float64
	obList []Observer //注册链表
}

func NewShareNotifier(price float64) *ShareNotifier {
	return &ShareNotifier{Price: price}
}

func (share *ShareNotifier) Register(observer Observer) {
	share.obList = append(share.obList, observer)
}

func (share *ShareNotifier) Remove(observer Observer) {
	for i, ob := range share.obList {
		if ob == observer {
			share.obList = append(share.obList[:i], share.obList[i+1:]...)
		}
	}
}

func (share *ShareNotifier) Notify(event Event) {
	for _, ob := range share.obList {
		ob.Receive(event)
	}
}
