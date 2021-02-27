package patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {

	in := &invoker{}

	cmda := &command{kind: 1}
	cmdb := &command{kind: 2}
	in.addCommand(cmda, cmdb)

	ra := &receiver{id: 10001}
	rb := &receiver{id: 10002}
	rc := &receiver{id: 10003}

	cmda.addReceiver(ra, rb, rc)
	cmdb.addReceiver(rb)

	in.excuseCommand()

	time.Sleep(1000)
	fmt.Println("-------------------------------")

	in.delCommand(cmda)
	in.excuseCommand()

}
