package domain

import "fmt"

type ItemType string

const (
	ItemTypeArmor  ItemType = "armor"
	ItemTypeWeapon ItemType = "weapon"
	ItemTypePotion ItemType = "potion"
)

type Item struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Type        ItemType `json:"type"`
	Description string   `json:"description,omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf(
		"Item: id: %s, name: %s, type: %s, description: %s",
		i.Id, i.Name, i.Type, i.Description)
}

type Equipment struct {
	Item
	Equipped bool `json:"equipped"`
}

func (e Equipment) String() string {
	return fmt.Sprintf("Equipment: equipped: %t", e.Equipped)
}
