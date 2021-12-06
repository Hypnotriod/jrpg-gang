package domain

import "fmt"

type ItemType string

const (
	ItemTypeArmor      ItemType = "armor"
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeDisposable ItemType = "disposable"
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
