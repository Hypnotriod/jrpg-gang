package test

import (
	"fmt"
	"testing"
)

func TestPrintArmor(t *testing.T) {
	fmt.Println(getBaseGloves(t))
}

func TestPrintWeapon(t *testing.T) {
	fmt.Println(getBaseSword(t))
}
