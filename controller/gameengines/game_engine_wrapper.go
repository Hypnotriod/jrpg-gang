package gameengines

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"sync"
)

type GameEngineWrapper struct {
	sync.RWMutex
	nextPhaseTimer      engine.GameEngineTimer
	broadcastGameAction func(playerIds []engine.PlayerId, result *engine.GameEvent)
	engine              *engine.GameEngine
}

func NewGameEngineWrapper(
	engine *engine.GameEngine,
	broadcastGameAction func(playerIds []engine.PlayerId, result *engine.GameEvent)) *GameEngineWrapper {
	w := &GameEngineWrapper{}
	w.engine = engine
	w.broadcastGameAction = broadcastGameAction
	return w
}

func (w *GameEngineWrapper) Dispose() {
	w.engine.Dispose()
	w.stopNextPhaseTimer()
	w.broadcastGameAction = nil
}

func (w *GameEngineWrapper) ReadGameState(playerId engine.PlayerId) (*engine.GameEvent, []engine.PlayerId, bool) {
	event := w.engine.NewGameEventWithPlayersInfo()
	broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
}

func (w *GameEngineWrapper) ReadPlayerInfo(playerId engine.PlayerId) (engine.PlayerInfo, bool) {
	unit := w.engine.FindActorByPlayerId(playerId)
	if unit == nil || unit.PlayerInfo == nil {
		return engine.PlayerInfo{}, false
	}
	return *unit.PlayerInfo, true
}

func (w *GameEngineWrapper) ExecuteUserAction(action domain.Action, playerId engine.PlayerId) (*engine.GameEvent, []engine.PlayerId, bool) {
	event := w.engine.ExecuteUserAction(action, playerId)
	broadcastPlayerIds := []engine.PlayerId{}
	if event.Phase != event.NextPhase {
		w.setNextPhaseTimer()
	}
	if event.UnitActionResult.Result.Result == domain.ResultAccomplished {
		broadcastPlayerIds = w.engine.GetRestPlayerIds(playerId)
	}
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
}

func (w *GameEngineWrapper) ReadyForNextPhase(playerId engine.PlayerId, isReady bool) (*engine.GameEvent, []engine.PlayerId, bool) {
	if !w.engine.NextPhaseRequired() || w.engine.AllActorsDead() {
		return nil, nil, false
	}
	w.engine.UpdateActorReady(playerId, isReady)
	if w.engine.AllActorsReady() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
		return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
	}
	event := w.engine.NewGameEventWithPlayersInfo()
	broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
}

func (w *GameEngineWrapper) SkipToNextPhase() (*engine.GameEvent, []engine.PlayerId, bool) {
	if w.engine.AllActorsDead() {
		return nil, nil, false
	}
	w.engine.ClearActiveUnitActionPoints()
	event := w.engine.NextPhase()
	w.setNextPhaseTimer()
	broadcastPlayerIds := w.engine.GetPlayerIds()
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
}

func (w *GameEngineWrapper) ConnectionStatusChanged(playerId engine.PlayerId, isOffline bool) (*engine.GameEvent, []engine.PlayerId, bool) {
	w.engine.UpdateUserConnectionStatus(playerId, isOffline)
	broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
	event := w.engine.NewGameEventWithPlayersInfo()
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
}

func (w *GameEngineWrapper) LeaveGame(playerId engine.PlayerId) (*engine.GameEvent, []engine.PlayerId, domain.Unit, bool) {
	var unit domain.Unit
	if u := w.engine.FindActorByPlayerId(playerId); u != nil {
		unit = u.Unit
		if !u.IsDead {
			share := w.engine.TakeAShare()
			unit.Booty.Accumulate(share)
		}
	}
	w.engine.RemoveActor(playerId)
	if w.engine.NextPhaseRequired() && w.engine.AllActorsReady() && !w.engine.AllActorsDead() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
		return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, unit, true
	}
	playerIds := w.engine.GetPlayerIds()
	event := w.engine.NewGameEventWithPlayersInfo()
	if len(playerIds) == 0 {
		w.Dispose()
	}
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), playerIds, unit, true
}

func (w *GameEngineWrapper) RemoveUser(playerId engine.PlayerId) (*engine.GameEvent, []engine.PlayerId, bool) {
	w.engine.RemoveActor(playerId)
	playerIds := w.engine.GetPlayerIds()
	if len(playerIds) == 0 {
		w.Dispose()
		return nil, nil, false
	}
	if w.engine.NextPhaseRequired() && w.engine.AllActorsReady() && !w.engine.AllActorsDead() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastPlayerIds := w.engine.GetRestPlayerIds(playerId)
		return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), broadcastPlayerIds, true
	}
	event := w.engine.NewGameEventWithPlayersInfo()
	return event.WithPhaseTimeout(w.nextPhaseTimer.SecondsLeft()), playerIds, true
}
