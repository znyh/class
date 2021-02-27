package patterns

import (
	"fmt"
)

type idepartment interface {
	sentmessage(string)
	getmessage(string)
}

type media struct {
	technology
	market
}

func (m *media) forwardmessage(i idepartment, msg string) {
	switch i.(type) {
	case *technology:
		m.market.getmessage(msg)
	case *market:
		m.technology.getmessage(msg)
	}
}

type technology struct {
	*media
}

func (t *technology) sentmessage(msg string) {
	t.media.forwardmessage(t, msg)
}
func (t *technology) getmessage(msg string) {
	fmt.Printf("技术部收到消息: %s\n", msg)
}

type market struct {
	*media
}

func (m *market) sentmessage(msg string) {
	m.media.forwardmessage(m, msg)
}
func (m *market) getmessage(msg string) {
	fmt.Printf("市场部部收到消息: %s\n", msg)
}
