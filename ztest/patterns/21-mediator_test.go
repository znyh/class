package patterns

import (
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := &media{}

	technical := technology{mediator}
	market := market{mediator}

	mediator.market = market
	mediator.technology = technical

	technical.sentmessage("开发已经完成")
	market.sentmessage("市场推广中")
}
