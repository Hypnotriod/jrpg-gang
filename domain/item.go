package domain

type Item struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Equipment struct {
	Item
	Equipped bool `json:"equipped"`
}
