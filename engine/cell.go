package engine

type CellConfiguration struct {
	Position Position `json:"position"`
}

type Cell struct {
	CellConfiguration
	Unit *GameUnit `json:"unit,omitempty"`
}
