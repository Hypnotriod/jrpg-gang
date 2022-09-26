package domain

import (
	"fmt"
	"jrpg-gang/util"
)

type UnitBooty struct {
	Coins int `json:"coins"`
	Ruby  int `json:"ruby,omitempty"`
}

func (b UnitBooty) String() string {
	return fmt.Sprintf("coins: %d, ruby: %d",
		b.Coins, b.Ruby)
}

func (b *UnitBooty) Accumulate(booty UnitBooty) {
	b.Coins += booty.Coins
	b.Ruby += booty.Ruby
}

func (b *UnitBooty) Reduce(booty UnitBooty, quantity uint) {
	b.Coins -= booty.Coins * int(quantity)
	b.Ruby -= booty.Ruby * int(quantity)
}

func (b *UnitBooty) TakeAShare(participants int) UnitBooty {
	share := UnitBooty{}
	share.Coins = b.Coins / participants
	share.Ruby = b.Ruby / participants
	b.Reduce(share, 1)
	return share
}

func (b *UnitBooty) Normalize() {
	b.Coins = util.Max(b.Coins, 0)
	b.Ruby = util.Max(b.Ruby, 0)
}

func (b *UnitBooty) Check(booty UnitBooty, quantity uint) bool {
	return b.Coins*int(quantity) <= booty.Coins &&
		b.Ruby*int(quantity) <= booty.Ruby
}
