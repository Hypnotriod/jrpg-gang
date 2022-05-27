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

func (b *UnitBooty) Reduce(booty UnitBooty) {
	b.Coins += booty.Coins
	b.Ruby += booty.Ruby
}

func (b *UnitBooty) Normalize() {
	b.Coins = util.Max(b.Coins, 0)
	b.Ruby = util.Max(b.Ruby, 0)
}

func (b *UnitBooty) Check(booty UnitBooty) bool {
	return b.Coins >= booty.Coins &&
		b.Ruby >= booty.Ruby
}
