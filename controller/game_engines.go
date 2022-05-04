package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameEngines struct {
	sync.RWMutex
	rndGen  *util.RndGen
	engines []*engine.GameEngine
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	return e
}

func (e *GameEngines) Add(userUids []uint, engine *engine.GameEngine) {

}
