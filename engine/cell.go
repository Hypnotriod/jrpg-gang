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
	Code     string            `json:"code,omitempty"`
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"faction ids: %v, type: %s",
		c.Factions,
		c.Type,
	)
}

func (c *Cell) Clone() *Cell {
	r := &Cell{}
	r.Factions = append(r.Factions, c.Factions...)
	r.Type = c.Type
	r.Code = c.Code
	return r
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
