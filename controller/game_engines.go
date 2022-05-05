package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameEngines struct {
	sync.RWMutex
	rndGen         *util.RndGen
	engines        []*engine.GameEngine
	userIdToEngine map[engine.UserId]*engine.GameEngine
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	e.userIdToEngine = make(map[engine.UserId]*engine.GameEngine)
	return e
}

func (e *GameEngines) Add(engine *engine.GameEngine) {
	defer e.Unlock()
	e.Lock()
	e.engines = append(e.engines, engine)
	for _, userId := range engine.GetUserIds() {
		e.userIdToEngine[userId] = engine
	}
}

func (e *GameEngines) HasUser(userId engine.UserId) bool {
	defer e.RUnlock()
	e.RLock()
	_, ok := e.userIdToEngine[userId]
	return ok
}
