package engine

import "fmt"

type CellType string

const (
	CellTypeSpace    = "space"
	CellTypeObstacle = "obstacle"
)

type Cell struct {
	FractionIds []uint   `json:"fractionIds"`
	Type        CellType `json:"type"`
	Kind        string   `json:"kind,omitempty"`
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"fraction id: %v, type: %s",
		c.FractionIds,
		c.Type,
	)
}

func (c *Cell) CanPlaceUnit() bool {
	return c.Type == CellTypeSpace
}

func (c *Cell) ContainsFractionId(fractionId uint) bool {
	for _, v := range c.FractionIds {
		if v == fractionId {
			return true
		}
	}
	return false
}
