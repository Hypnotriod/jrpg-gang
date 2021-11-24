package test

import (
	"fmt"
	"testing"
)

func TestAttackUnit(t *testing.T) {
	baseSword := getBaseSword(t)
	northTroll := getNorthTroll(t)
	weakLuckyHero := getWeakLuckyHero(t)
	weakLuckyHero.Items = append(northTroll.Items, baseSword)
	baseSword.Equipped = true
	fmt.Println(weakLuckyHero)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	damage, tempImpact, success := weakLuckyHero.Attack(northTroll, baseSword.Impact)
	fmt.Printf("Attack with damage: {%v}, impact: %v, success: %t\n", damage, tempImpact, success)
	fmt.Println()
	fmt.Println(northTroll)
	fmt.Println()
	fmt.Println(weakLuckyHero)
	fmt.Println()
}
