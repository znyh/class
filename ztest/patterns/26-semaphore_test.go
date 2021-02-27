package patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {

	se := newsemaphore(1, time.Second)

	if err := se.acquire(); err != nil {
		t.Errorf("bad semaphore,se.acquire() err:%s\n", err)
	}

	if err := se.release(); err != nil {
		t.Errorf("bad semaphore,se.release() err:%s\n", err)
	}

	if err := se.release(); err == nil {
		t.Errorf("bad semaphore,se.release2() err:%s\n", err)
	}
}

func TestSemaphore2(t *testing.T) {

	se := newsemaphore(0, time.Second)

	done := make(chan bool)

	go func() {
		done <- true
		if err := se.acquire(); err != nil {
			t.Errorf("bad semaphore,se.acquire() err:%s\n", err)
		}
	}()

	go func() {
		if err := se.release(); err != nil {
			t.Errorf("bad semaphore,se.release() err:%s\n", err)
		}
		<-done
	}()

	fmt.Println("all done")
}
