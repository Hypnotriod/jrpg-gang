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

type Item struct {
	Uid         uint     `json:"uid,omitempty"`
	Name        string   `json:"name"`
	Type        ItemType `json:"type"`
	Description string   `json:"description,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf(
		"name: %s, uid: %d, type: %s, description: %s",
		i.Name, i.Uid, i.Type, i.Description)
}
