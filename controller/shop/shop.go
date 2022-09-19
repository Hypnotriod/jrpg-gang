package shop

import (
	"jrpg-gang/controller/factory"
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"sync"
)

type Shop struct {
	mu   sync.RWMutex
	shop *engine.GameShop
}

func NewShop() *Shop {
	s := &Shop{}
	s.shop = engine.NewGameShop(factory.NewTestShopItems())
	return s
}

func (s *Shop) GetStatus() engine.GameShop {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return *s.shop
}

func (s *Shop) ExecuteAction(action domain.Action, user *users.User) *domain.ActionResult {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return s.shop.ExecuteAction(action, &user.Unit.Unit, user.RndGen)
}
