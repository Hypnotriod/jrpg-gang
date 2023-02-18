package gameengines

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameEngines struct {
	mu               sync.RWMutex
	rndGen           *util.RndGen
	playerIdToEngine map[engine.PlayerId]*GameEngineWrapper
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	e.playerIdToEngine = make(map[engine.PlayerId]*GameEngineWrapper)
	return e
}

func (e *GameEngines) Register(wrapper *GameEngineWrapper) (*engine.GameEvent, []engine.PlayerId) {
	defer e.mu.Unlock()
	e.mu.Lock()
	for _, playerId := range wrapper.engine.GetPlayerIds() {
		e.playerIdToEngine[playerId] = wrapper
	}
	wrapper.setNextPhaseTimer()
	timeout := wrapper.nextPhaseTimer.SecondsLeft()
	playerIds := wrapper.engine.GetPlayerIds()
	state := wrapper.engine.NewGameEventWithPlayersInfo()
	return state.WithPhaseTimeout(timeout), playerIds
}

func (e *GameEngines) Find(playerId engine.PlayerId) (*GameEngineWrapper, bool) {
	e.mu.RLock()
	wrapper, ok := e.playerIdToEngine[playerId]
	e.mu.RUnlock()
	return wrapper, ok
}

func (e *GameEngines) Unregister(playerId engine.PlayerId) (*GameEngineWrapper, bool) {
	e.mu.Lock()
	wrapper, ok := e.playerIdToEngine[playerId]
	if !ok {
		e.mu.Unlock()
		return wrapper, ok
	}
	delete(e.playerIdToEngine, playerId)
	e.mu.Unlock()
	return wrapper, ok
}

func (e *GameEngines) IsUserInGame(playerId engine.PlayerId) bool {
	defer e.mu.RUnlock()
	e.mu.RLock()
	_, ok := e.playerIdToEngine[playerId]
	return ok
}
