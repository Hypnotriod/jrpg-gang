package domain

import "fmt"

type ItemType string

const (
	ItemTypeArmor      ItemType = "armor"
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeMagic      ItemType = "magic"
	ItemTypeDisposable ItemType = "disposable"
	ItemTypeAmmunition ItemType = "ammunition"
	ItemTypeNone       ItemType = "none"
)

type ItemCode string

const ItemCodeEmpty ItemCode = ""

type Item struct {
	Uid         uint      `json:"uid,omitempty"`
	Code        ItemCode  `json:"code"`
	Name        string    `json:"name"`
	Type        ItemType  `json:"type"`
	Price       UnitBooty `json:"price,"`
	Description string    `json:"description,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf(
		"name: %s, code: %s, uid: %d, type: %s, description: %s",
		i.Name, i.Code, i.Uid, i.Type, i.Description)
}
