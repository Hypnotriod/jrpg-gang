package gameengines

import (
	"time"
)

func (w *GameEngineWrapper) setNextPhaseTimer() {
	w.stopNextPhaseTimer()
	if w.engine.AllActorsDead() {
		return
	}
	timerId := w.nextPhaseTimerId
	w.nextPhaseTimer = time.AfterFunc(time.Duration(5)*time.Second, func() { //todo configure time
		defer w.Unlock()
		w.Lock()
		if timerId != w.nextPhaseTimerId {
			return
		}
		if result, broadcastUserIds, ok := w.NextPhase(); ok {
			w.broadcastGameAction(broadcastUserIds, result)
		}
	})
}

func (w *GameEngineWrapper) stopNextPhaseTimer() {
	if w.nextPhaseTimer != nil {
		w.nextPhaseTimer.Stop()
		w.nextPhaseTimer = nil
	}
	w.nextPhaseTimerId++
}
