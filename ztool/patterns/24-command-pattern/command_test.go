package command

import "testing"

var (
	kind TYPE = "a"
	mold TYPE = "b"
)

func TestInvoker_ExecuteCommand(t *testing.T) {

	receivera := &ReceiverA{}
	receiverb := &ReceiverB{}

	commanda := CreateCommand(kind, receivera)
	conmandb := CreateCommand(mold, receiverb)

	invoker := new(Invoker)
	invoker.AddCommand(commanda)
	invoker.AddCommand(conmandb)

	invoker.ExecuteCommand() //调用： 接收者a执行A的操作，b执行B的操作
}
