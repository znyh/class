package patterns

import (
	"container/list"
	"fmt"
)

//// receive..receive..
type ireceive interface {
	excuse(int)
}

type receiver struct {
	id int
}

func (r *receiver) excuse(commandkind int) {
	fmt.Printf("接收者:[id:%d],处理请求命令:[%d]\n", r.id, commandkind)
}

//// command..command..
type icommand interface {
	call()
}

type command struct {
	kind      int
	receivers *list.List
}

func (c *command) call() {
	for e := c.receivers.Front(); e != nil; e = e.Next() {
		e.Value.(*receiver).excuse(c.kind)
	}
}

func (c *command) addReceiver(rs ...ireceive) {
	if c.receivers == nil {
		c.receivers = list.New()
	}
	for _, r := range rs {
		c.receivers.PushBack(r)
	}
}

func (c *command) delReceiver(rs ...ireceive) {
	if c.receivers.Len() <= 0 {
		return
	}
	for _, r := range rs {
		for e := c.receivers.Front(); e != nil; e = e.Next() {
			if e.Value == r {
				c.receivers.Remove(e)
			}
		}
	}
}

//// invoker..invoker..
type invoker struct {
	commands *list.List
}

func (in *invoker) addCommand(cs ...icommand) {
	if in.commands == nil {
		in.commands = list.New()
	}
	for _, c := range cs {
		in.commands.PushBack(c)
	}
}

func (in *invoker) delCommand(cs ...icommand) {
	if in.commands.Len() <= 0 {
		return
	}
	for _, c := range cs {
		for e := in.commands.Front(); e != nil; e = e.Next() {
			if e.Value == c {
				in.commands.Remove(e)
			}
		}
	}
}

func (in *invoker) excuseCommand() {
	for e := in.commands.Front(); e != nil; e = e.Next() {
		e.Value.(*command).call()
	}
}
