package mediator

import "testing"

func TestMediator_ForwardMessage(t *testing.T) {

	mediator := &Mediator{}

	technical := Technical{mediator}
	market := Market{mediator}

	mediator.Market = market
	mediator.Technical = technical

	technical.SendMess("开发已经完成")
	market.SendMess("市场推广中")
}
