package gameengines

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameEngines struct {
	mu             sync.RWMutex
	rndGen         *util.RndGen
	userIdToEngine map[engine.UserId]*GameEngineWrapper
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	e.userIdToEngine = make(map[engine.UserId]*GameEngineWrapper)
	return e
}

func (e *GameEngines) Register(wrapper *GameEngineWrapper) {
	e.mu.Lock()
	for _, userId := range wrapper.engine.GetUserIds() {
		e.userIdToEngine[userId] = wrapper
	}
	e.mu.Unlock()
}

func (e *GameEngines) Find(userId engine.UserId) (*GameEngineWrapper, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	return wrapper, ok
}

func (e *GameEngines) Unregister(userId engine.UserId) (*GameEngineWrapper, bool) {
	e.mu.Lock()
	wrapper, ok := e.userIdToEngine[userId]
	if !ok {
		e.mu.Unlock()
		return wrapper, ok
	}
	delete(e.userIdToEngine, userId)
	e.mu.Unlock()
	return wrapper, ok
}

func (e *GameEngines) IsUserInGame(userId engine.UserId) bool {
	defer e.mu.RUnlock()
	e.mu.RLock()
	_, ok := e.userIdToEngine[userId]
	return ok
}
