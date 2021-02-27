package patterns

import (
	"strings"
)

type iexpress interface {
	interpretate() bool
}

type content struct {
	val string
}

func (c *content) getval() string {
	return c.val
}

func createExpress(kind int, left, right content) iexpress {
	switch kind {
	case 1:
		return &equal{left: left, right: right}
	case 2:
		return &contain{left: left, right: right}
	default:
		return nil
	}
}

type equal struct {
	left  content
	right content
}

func (e *equal) interpretate() bool {
	return e.right.getval() == e.left.getval()
}

type contain struct {
	left  content
	right content
}

func (c *contain) interpretate() bool {
	return strings.Contains(c.left.getval(), c.right.getval())
}
