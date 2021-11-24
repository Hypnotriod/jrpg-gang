package test

import (
	"fmt"
	"testing"
)

func TestAttackUnit(t *testing.T) {
	baseSword := getBaseSword(t)
	northTroll := getNorthTroll(t)
	weakHero := getWeakHero(t)
	weakHero.Items = append(northTroll.Items, baseSword)
	baseSword.Equipped = true
	damage, success := weakHero.Attack(northTroll, baseSword.Impact)
	fmt.Println(weakHero)
	fmt.Println(northTroll)
	fmt.Printf("damage: %v, success: %t\n", damage, success)
	fmt.Println(northTroll)
}
