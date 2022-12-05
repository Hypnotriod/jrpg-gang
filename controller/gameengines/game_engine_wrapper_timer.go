package gameengines

import (
	"jrpg-gang/engine"
	"time"
)

const NEXT_PHASE_TIMEOUT_SHORT_SEC int = 5
const NEXT_PHASE_TIMEOUT_MEDIUM_SEC int = 32
const NEXT_PHASE_TIMEOUT_LONG_SEC int = 62

func (w *GameEngineWrapper) getNextPhaseTimeout() (int, bool) {
	switch w.engine.GetPhase() {
	case engine.GamePhasePrepareUnit,
		engine.GamePhaseBattleComplete:
		return NEXT_PHASE_TIMEOUT_LONG_SEC, true
	case engine.GamePhaseMakeMoveOrAction,
		engine.GamePhaseMakeAction:
		return NEXT_PHASE_TIMEOUT_MEDIUM_SEC, true
	case engine.GamePhaseReadyForStartRound,
		engine.GamePhaseMakeMoveOrActionAI,
		engine.GamePhaseMakeActionAI,
		engine.GamePhaseRetreatAction,
		engine.GamePhaseActionComplete:
		return NEXT_PHASE_TIMEOUT_SHORT_SEC, true
	}
	return 0, false
}

func (w *GameEngineWrapper) setNextPhaseTimer() {
	w.stopNextPhaseTimer()
	if w.engine.AllActorsDead() {
		return
	}
	timeout, ok := w.getNextPhaseTimeout()
	if !ok {
		return
	}
	timerId := w.nextPhaseTimerId
	w.nextPhaseTimer = time.AfterFunc(time.Duration(timeout)*time.Second, func() {
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
