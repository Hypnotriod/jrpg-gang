package test

import (
	"fmt"
	"jrpg-gang/util"
	"testing"
)

func TestAttackUnit(t *testing.T) {
	util.ApplyRandomSeed()
	baseSword := newBaseSword(t)
	northTroll := newNorthTroll(t)
	hero := newAgileHero(t)
	hero.Items = append(northTroll.Items, baseSword)
	baseSword.Equipped = true
	fmt.Println(hero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	damage, tempImpact := hero.Attack(northTroll, baseSword.Impact)
	fmt.Printf("Attack with damage: %v, impact: %v\n", damage, tempImpact)
	fmt.Println()
	fmt.Println(northTroll)
	for len(northTroll.Impact) > 0 {
		fmt.Println()
		fmt.Println("Turn Passed")
		fmt.Println()
		northTroll.ApplyDamageImpactOnNextTurn()
		northTroll.ReduceEnhancementOnNextTurn()
		fmt.Println(northTroll)
	}
}
