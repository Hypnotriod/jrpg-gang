package gameengines

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"sync"
	"time"
)

type GameEngineWrapper struct {
	sync.RWMutex
	nextPhaseTimer      *time.Timer
	nextPhaseTimerId    uint32
	broadcastGameAction func(userIds []engine.UserId, result *engine.GameEvent)
	engine              *engine.GameEngine
}

func NewGameEngineWrapper(
	engine *engine.GameEngine,
	broadcastGameAction func(userIds []engine.UserId, result *engine.GameEvent)) *GameEngineWrapper {
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

func (w *GameEngineWrapper) ReadGameState(userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	event := w.engine.NewGameEventWithPlayersInfo()
	broadcastUserIds := w.engine.GetRestUserIds(userId)
	return event, broadcastUserIds, true
}

func (w *GameEngineWrapper) ReadPlayerInfo(userId engine.UserId) (engine.PlayerInfo, bool) {
	unit := w.engine.FindActorByUserId(userId)
	if unit == nil || unit.PlayerInfo == nil {
		return engine.PlayerInfo{}, false
	}
	return *unit.PlayerInfo, true
}

func (w *GameEngineWrapper) ExecuteUserAction(action domain.Action, userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	event := w.engine.ExecuteUserAction(action, userId)
	broadcastUserIds := []engine.UserId{}
	if event.Phase != event.NextPhase {
		w.setNextPhaseTimer()
	}
	if event.UnitActionResult.Result.Result == domain.ResultAccomplished {
		broadcastUserIds = w.engine.GetRestUserIds(userId)
	}
	return event, broadcastUserIds, true
}

func (w *GameEngineWrapper) ReadyForNextPhase(userId engine.UserId, isReady bool) (*engine.GameEvent, []engine.UserId, bool) {
	if !w.engine.NextPhaseRequired() || w.engine.AllActorsDead() {
		return nil, nil, false
	}
	w.engine.UpdateActorReady(userId, isReady)
	if w.engine.AllActorsReady() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastUserIds := w.engine.GetRestUserIds(userId)
		return event, broadcastUserIds, true
	}
	event := w.engine.NewGameEventWithPlayersInfo()
	broadcastUserIds := w.engine.GetRestUserIds(userId)
	return event, broadcastUserIds, true
}

func (w *GameEngineWrapper) NextPhase() (*engine.GameEvent, []engine.UserId, bool) {
	if w.engine.AllActorsDead() {
		return nil, nil, false
	}
	event := w.engine.NextPhase()
	w.setNextPhaseTimer()
	broadcastUserIds := w.engine.GetUserIds()
	return event, broadcastUserIds, true
}

func (w *GameEngineWrapper) ConnectionStatusChanged(userId engine.UserId, isOffline bool) (*engine.GameEvent, []engine.UserId, bool) {
	w.engine.UpdateUserConnectionStatus(userId, isOffline)
	broadcastUserIds := w.engine.GetRestUserIds(userId)
	event := w.engine.NewGameEventWithPlayersInfo()
	return event, broadcastUserIds, true
}

func (w *GameEngineWrapper) LeaveGame(userId engine.UserId) (*engine.GameEvent, []engine.UserId, domain.Unit, bool) {
	var unit domain.Unit
	if u := w.engine.FindActorByUserId(userId); u != nil {
		unit = u.Unit
		if !u.IsDead {
			share := w.engine.TakeAShare()
			unit.Booty.Accumulate(share)
		}
	}
	w.engine.RemoveActor(userId)
	if w.engine.NextPhaseRequired() && w.engine.AllActorsReady() && !w.engine.AllActorsDead() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastUserIds := w.engine.GetRestUserIds(userId)
		return event, broadcastUserIds, unit, true
	}
	userIds := w.engine.GetUserIds()
	event := w.engine.NewGameEventWithPlayersInfo()
	if len(userIds) == 0 {
		w.Dispose()
	}
	return event, userIds, unit, true
}

func (w *GameEngineWrapper) RemoveUser(userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	w.engine.RemoveActor(userId)
	userIds := w.engine.GetUserIds()
	if len(userIds) == 0 {
		w.Dispose()
		return nil, nil, false
	}
	if w.engine.NextPhaseRequired() && w.engine.AllActorsReady() && !w.engine.AllActorsDead() {
		event := w.engine.NextPhase()
		w.setNextPhaseTimer()
		broadcastUserIds := w.engine.GetRestUserIds(userId)
		return event, broadcastUserIds, true
	}
	event := w.engine.NewGameEventWithPlayersInfo()
	return event, userIds, true
}
