package domain

import (
	"jrpg-gang/util"
)

type UnitBooty struct {
	Coins int `json:"coins"`
	Ruby  int `json:"ruby,omitempty"`
	W     int `json:"weight,omitempty"`
}

func (b UnitBooty) Weight() int {
	return b.W
}

func (b *UnitBooty) Accumulate(booty UnitBooty) {
	b.Coins += booty.Coins
	b.Ruby += booty.Ruby
}

func (b *UnitBooty) Reduce(booty UnitBooty, quantity uint) {
	b.Coins -= booty.Coins * int(quantity)
	b.Ruby -= booty.Ruby * int(quantity)
}

func (b *UnitBooty) MultiplyAll(factor float32) {
	b.Coins = int(util.MultiplyWithRounding(float32(b.Coins), factor))
	b.Ruby = int(util.MultiplyWithRounding(float32(b.Ruby), factor))
}

func (b *UnitBooty) TakeAShare(participants int) UnitBooty {
	share := UnitBooty{
		Coins: int(util.Ceil(float32(b.Coins) / float32(participants))),
		Ruby:  int(util.Ceil(float32(b.Ruby) / float32(participants))),
	}
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
