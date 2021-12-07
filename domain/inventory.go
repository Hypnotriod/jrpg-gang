package domain

import "fmt"

type Inventory struct {
	Weapon     []Weapon     `json:"weapon,omitempty"`
	Armor      []Armor      `json:"armor,omitempty"`
	Disposable []Disposable `json:"disposable,omitempty"`
}

func (i Inventory) String() string {
	return fmt.Sprintf(
		"weapon: %v, armor: %v, disposable: %v",
		i.Weapon, i.Armor, i.Disposable)
}
