package test

import (
	"fmt"
	"testing"
)

func TestPrintArmor(t *testing.T) {
	fmt.Println(newBaseGloves(t))
}

func TestPrintWeapon(t *testing.T) {
	fmt.Println(newBaseSword(t))
}
