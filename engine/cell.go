package engine

import "fmt"

type CellType string

const (
	CellTypeSpace    = "space"
	CellTypeObstacle = "obstacle"
)

type Cell struct {
	FractionId uint
	Type       CellType
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"fraction id: %d, type: %s",
		c.FractionId,
		c.Type,
	)
}

func (c *Cell) CanPlaceUnit() bool {
	return c.Type == CellTypeSpace
}
