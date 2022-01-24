package engine

type CellType string

const (
	CellTypeSpace    = "space"
	CellTypeObstacle = "obstacle"
)

type Cell struct {
	FractionId uint
	Type       CellType
}
