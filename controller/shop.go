package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"sync"
)

type Shop struct {
	sync.RWMutex
	shop *engine.GameShop
}

func NewShop() *Shop {
	s := &Shop{}
	s.shop = engine.NewGameShop(NewTestShopItems())
	return s
}

func (s *Shop) GetStatus() engine.GameShop {
	defer s.RUnlock()
	s.RLock()
	return *s.shop
}

func (s *Shop) ExecuteAction(action domain.Action, user *User) *domain.ActionResult {
	defer s.RUnlock()
	s.RLock()
	return s.shop.ExecuteAction(action, &user.unit.Unit, user.rndGen)
}
