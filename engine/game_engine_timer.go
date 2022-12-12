package engine

import (
	"jrpg-gang/util"
	"time"
)

type GameEngineTimer struct {
	id        int
	duration  time.Duration
	startTime time.Time
	timer     *time.Timer
}

func (t *GameEngineTimer) AfterFunc(d time.Duration, f func()) int {
	if t.timer != nil {
		t.timer.Stop()
	}
	t.duration = d
	t.timer = time.AfterFunc(d, f)
	t.startTime = time.Now()
	return t.id
}

func (t *GameEngineTimer) Stop() bool {
	defer func() {
		t.timer = nil
	}()
	t.id++
	if t.timer != nil {
		return t.timer.Stop()
	}
	return false
}

func (t *GameEngineTimer) SecondsLeft() float32 {
	if t.timer == nil {
		return 0
	}
	result := (t.duration - time.Since(t.startTime)).Seconds()
	result = util.Round(result)
	result = util.Max(result, 0)
	return float32(result)
}

func (t *GameEngineTimer) Id() int {
	return t.id
}
