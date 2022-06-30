package util

import (
	"time"
)

type Timer struct {
	timer  *time.Timer
	cancel chan struct{}
}

func NewTimer(duration time.Duration, complete func()) *Timer {
	t := &Timer{}
	t.timer = time.NewTimer(duration)
	t.cancel = make(chan struct{})
	go func() {
		select {
		case <-t.timer.C:
			complete()
		case <-t.cancel:
		}
	}()
	return t
}

func NewTimerWithCancel(duration time.Duration, complete func(), cancel func()) *Timer {
	t := &Timer{}
	t.timer = time.NewTimer(duration)
	t.cancel = make(chan struct{})
	go func() {
		select {
		case <-t.timer.C:
			complete()
		case <-t.cancel:
			cancel()
		}
	}()
	return t
}

func (t *Timer) Cancel() {
	t.timer.Stop()
	t.cancel <- struct{}{}
}
