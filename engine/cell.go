package engine

type CellType int

const (
	CellTypeSpace    CellType = 0
	CellTypeObstacle CellType = 1
)

type Cell struct {
	Factions []GameUnitFaction `json:"factions"`
	Type     CellType          `json:"type,omitempty"`
	Code     string            `json:"code,omitempty"`
}

func (c *Cell) Clone() *Cell {
	r := &Cell{}
	r.Factions = []GameUnitFaction{}
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
