package semaphore

import (
	"fmt"
	"testing"
	"time"
)

func TestNoRelease(t *testing.T) {
	ticket, timeout := 1, 2*time.Second
	sem := New(ticket, timeout)

	if err := sem.Acquire(); err != nil {
		fmt.Println("sem.Acquire()", err)
	}

	if err := sem.Release(); err != nil {
		fmt.Println("sem.Release()_1 :", err)
	}
	if err := sem.Release(); err != nil {
		fmt.Println("sem.Release()_2 :", err)
	}
}

func TestNoTicket(t *testing.T) {
	ticket, timeout := 0, 1*time.Second
	sem := New(ticket, timeout)

	if err := sem.Acquire(); err != nil {
		if err == ErrNoTickets {
			fmt.Println(err)
		}
	}
}
