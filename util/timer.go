package util

import (
	"sync"
	"time"
)

type Timer struct {
	mu        sync.Mutex
	timer     *time.Timer
	cancel    chan struct{}
	cancelled bool
	completed bool
}

func NewTimer(duration time.Duration, complete func()) *Timer {
	t := &Timer{}
	t.timer = time.NewTimer(duration)
	t.cancel = make(chan struct{})
	go t.wait(complete, func() {})
	return t
}

func NewTimerWithCancel(duration time.Duration, complete func(), cancel func()) *Timer {
	t := &Timer{}
	t.timer = time.NewTimer(duration)
	t.cancel = make(chan struct{})
	go t.wait(complete, cancel)
	return t
}

func (t *Timer) Cancel() {
	t.mu.Lock()
	if t.completed {
		t.mu.Unlock()
		return
	}
	t.cancelled = true
	t.mu.Unlock()
	t.timer.Stop()
	t.cancel <- struct{}{}
}

func (t *Timer) wait(complete func(), cancel func()) {
	for {
		select {
		case <-t.timer.C:
			t.mu.Lock()
			if !t.cancelled {
				t.completed = true
				t.mu.Unlock()
				complete()
				return
			}
			t.mu.Unlock()
		case <-t.cancel:
			cancel()
			return
		}
	}
}
