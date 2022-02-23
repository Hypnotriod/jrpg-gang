package engine

import "fmt"

type CellType string

const (
	CellTypeSpace    = "space"
	CellTypeObstacle = "obstacle"
)

type Cell struct {
	Factions []GameUnitFaction `json:"factions"`
	Type     CellType          `json:"type"`
	Name     string            `json:"name,omitempty"`
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"faction ids: %v, type: %s",
		c.Factions,
		c.Type,
	)
}

func (c *Cell) CanPlaceUnit() bool {
	return c.Type == CellTypeSpace
}

func (c *Cell) ContainsFaction(faction GameUnitFaction) bool {
	for _, v := range c.Factions {
		if v == faction {
			return true
		}
	}
	return false
}
