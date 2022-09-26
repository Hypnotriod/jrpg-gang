package domain

import (
	"fmt"
	"testing"
)

func TestAttackUnitBySword(t *testing.T) {
	hero := newAgileHeroWithWeapon(t)
	northTroll := newNorthTroll(t)
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	hero.Equip(4000)
	actionResult := hero.UseInventoryItemOnTarget(northTroll, 4000)
	fmt.Printf("%s attacks %s with: {%v}\n", hero.Name, northTroll.Name, actionResult)
	fmt.Println()
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	for len(northTroll.Damage) > 0 {
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		northTroll.ApplyDamageOnNextTurn()
		northTroll.ReduceModificationOnNextTurn()
		fmt.Println(northTroll)
	}
}

func TestAttackUnitWithBow(t *testing.T) {
	hero := newAgileHeroWithWeapon(t)
	northTroll := newNorthTroll(t)
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	hero.Equip(4001)
	hero.Equip(5000)
	actionResult := hero.UseInventoryItemOnTarget(northTroll, 4001)
	fmt.Printf("%s attacks %s with: {%v}\n", hero.Name, northTroll.Name, actionResult)
	fmt.Println()
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	for len(northTroll.Damage) > 0 {
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		northTroll.ApplyDamageOnNextTurn()
		northTroll.ReduceModificationOnNextTurn()
		fmt.Println(northTroll)
	}
}
