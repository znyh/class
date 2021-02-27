package patterns

import (
	"sync"
)

var (
	ins  *single
	once sync.Once
)

type single struct {
}

func getSingleton() *single {
	once.Do(func() {
		ins = new(single)
	})
	return ins
}
