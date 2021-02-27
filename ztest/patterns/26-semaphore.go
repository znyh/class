package patterns

import (
	"errors"
	"fmt"
	"time"
)

type isemaphore interface {
	acquire() error
	release() error
}

type semaphore struct {
	ch      chan struct{}
	timeout time.Duration
}

func newsemaphore(cap int, timeout time.Duration) *semaphore {
	return &semaphore{
		ch:      make(chan struct{}, cap),
		timeout: timeout,
	}
}

func (s *semaphore) acquire() error {
	select {
	case s.ch <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return errors.New(fmt.Sprintf("acquire 超时，%v\n", s.timeout))
	}
}

func (s *semaphore) release() error {
	select {
	case <-s.ch:
		return nil
	case <-time.After(s.timeout):
		return errors.New(fmt.Sprintf("release 超时，%v\n", s.timeout))
	}
}
