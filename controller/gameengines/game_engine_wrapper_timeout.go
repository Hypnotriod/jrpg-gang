package gameengines

import (
	"jrpg-gang/engine"
	"time"
)

const NEXT_PHASE_TIMEOUT_SHORT_SEC int = 7
const NEXT_PHASE_TIMEOUT_MEDIUM_SEC int = 32
const NEXT_PHASE_TIMEOUT_LONG_SEC int = 62

func (w *GameEngineWrapper) getNextPhaseTimeout() (int, bool) {
	switch w.engine.GetPhase() {
	case engine.GamePhasePrepareUnit,
		engine.GamePhaseSpotComplete:
		return NEXT_PHASE_TIMEOUT_LONG_SEC, true
	case engine.GamePhaseMakeMoveOrAction,
		engine.GamePhaseMakeAction:
		return NEXT_PHASE_TIMEOUT_MEDIUM_SEC, true
	case engine.GamePhaseReadyForStartRound,
		engine.GamePhaseMakeMoveOrActionAI,
		engine.GamePhaseMakeActionAI,
		engine.GamePhaseRetreatAction,
		engine.GamePhaseActionComplete,
		engine.GamePhaseBeforeSpotComplete:
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
	timerId := w.nextPhaseTimer.Id()
	w.nextPhaseTimer.AfterFunc(time.Duration(timeout)*time.Second, func() {
		defer w.Unlock()
		w.Lock()
		if timerId != w.nextPhaseTimer.Id() {
			return
		}
		if result, broadcastPlayerIds, ok := w.SkipToNextPhase(); ok {
			w.broadcastGameAction(broadcastPlayerIds, result)
		}
	})
}

func (w *GameEngineWrapper) stopNextPhaseTimer() {
	w.nextPhaseTimer.Stop()
}
