package patterns

import (
	"fmt"
)

type reader struct {
}

func (r *reader) read() {
	fmt.Println("reader is reading")
}

type listenner struct {
}

func (l *listenner) listen() {
	fmt.Println("listenner is listenning")
}

type facade struct {
	r reader
	l listenner
}

func (f *facade) work() {
	f.r.read()
	f.l.listen()
}
