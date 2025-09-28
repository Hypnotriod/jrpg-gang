package domain

type ItemType string

const (
	ItemTypeArmor      ItemType = "armor"
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeMagic      ItemType = "magic"
	ItemTypeDisposable ItemType = "disposable"
	ItemTypeAmmunition ItemType = "ammunition"
	ItemTypeProvision  ItemType = "provision"
	ItemTypeNone       ItemType = "none"
)

type ItemCode string

const ItemCodeEmpty ItemCode = ""

type Item struct {
	Uid             uint      `json:"uid,omitempty"`
	Code            ItemCode  `json:"code"`
	Name            string    `json:"name"`
	Type            ItemType  `json:"type"`
	CanBeThrownAway bool      `json:"canBeThrownAway,omitempty"`
	CanBeSold       bool      `json:"canBeSold,omitempty"`
	Price           UnitBooty `json:"price,omitempty"`
	Description     string    `json:"description,omitempty"`
}
