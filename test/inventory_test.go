package test

import (
	"fmt"
	"testing"
)

func TestEquipItem(t *testing.T) {
	hero := newAgileHero(t)
	weapon := newBaseSword(t)
	gloves := newBaseGloves(t)

	fmt.Println(hero)
	fmt.Println()

	hero.Inventory.Add(weapon)
	hero.Inventory.Add(gloves)

	fmt.Println(hero)
	fmt.Println()

	hero.Equip(weapon.Id)
	hero.Equip(gloves.Id)

	fmt.Println(hero)
}
