package controller

import (
	"jrpg-gang/shop"
	"sync"
)

type Shop struct {
	sync.RWMutex
	shop *shop.GameShop
}

func NewShop() *Shop {
	s := &Shop{}
	s.shop = shop.NewGameShop(NewTestShopItems())
	return s
}

func (s *Shop) GetStatus() shop.GameShop {
	defer s.RUnlock()
	s.RLock()
	return *s.shop
}

func (s *Shop) Buy(user User, itemUid uint) bool {
	defer s.RUnlock()
	s.RLock()
	return s.shop.Buy(&user.unit.Unit, itemUid, user.rndGen)
}
