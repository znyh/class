package util

import (
	"container/list"
	"container/ring"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	a1 := []string{"a", "b", "c"}
	a2 := a1[1:2]
	log.Printf("a1:%+v a2:%+v\n", a1, a2)

	a2[0] = "bbb"
	log.Printf("a1:%+v a2:%+v\n", a1, a2)

	a1 = append(a1, "d")
	log.Printf("a1:%+v a2:%+v\n", a1, a2)

	a1 = []string(nil)
	log.Printf("a1:%+v a2:%+v\n", a1, a2)

}

func TestList(t *testing.T) {
	l := list.New()
	_a := l.PushBack("aaa")
	l.InsertAfter("bbb", _a)
	l.InsertBefore(0, _a)

	var st struct {
		id int
	}
	st.id = 1001

	l.PushBack(st)

	for e := l.Front(); e != nil; e = e.Next() {
		log.Printf("%+v", e.Value)
	}
	log.Printf("len:%d", l.Len())

	l.Remove(_a)

	for e := l.Back(); e != nil; e = e.Prev() {
		log.Printf("%+v", e.Value)
	}
	log.Printf("len:%d", l.Len())

	l.Init()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		log.Printf("%+v", e.Value)
	}
	log.Printf("len:%d", l.Len())

}

func TestRing(t *testing.T) {
	r := ring.New(5)
	for i := 0; i < 5; i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(i interface{}) {
		log.Printf("i:%d", i)
	})
}

func TestChannel(t *testing.T) {
	ch := make(chan int) //ch := make(chan int,2)
	close(ch)
	//ch <- 1
	//ch <- 2
	//ch <- 3
	log.Printf("==> %+v", <-ch)
}

func TestChannel1(t *testing.T) {
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		defer close(ch) //关闭后就不能往channel里放入数据
	}()

	for i := range ch {
		log.Printf("==> %+v", i)
	}
}

func TestChannel2(t *testing.T) {
	log.SetFlags(0)

	ready, done := make(chan struct{}), make(chan struct{})

	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}

	<-done
	<-done
	<-done

}

func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready

	log.Println("worker:", id, " start to process.")

	time.Sleep(time.Second * 1)

	log.Println("worker:", id, " finish its job.")

	done <- struct{}{}
}

func TestChannel3(t *testing.T) {
}

func merge(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	for _, in := range inputs {
		in := in
		go func(in <-chan uint64) {
			for {
				output <- <-in
			}
		}(in)
	}
	return output
}

func divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		out := out
		go func(out chan<- uint64) {
			for {
				out <- <-input
			}
		}(out)
	}
}

func TestChannel4(t *testing.T) {
	ch := make(chan int, 64) // 成果队列

	go producer4(3, ch) // 生成 3 的倍数的序列
	go producer4(5, ch) // 生成 5 的倍数的序列
	go consumer4(ch)    // 消费 生成的队列

	//time.Sleep(5 * time.Second)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func producer4(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func consumer4(in <-chan int) {
	for i := range in {
		log.Printf("i==> %d", i)
	}
}
