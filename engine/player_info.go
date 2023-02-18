package engine

type PlayerId string

const (
	PlayerIdEmpty PlayerId = ""
)

type PlayerInfo struct {
	Id        PlayerId      `json:"-"`
	Nickname  string        `json:"nickname"`
	Class     GameUnitClass `json:"class"`
	Level     uint          `json:"level"`
	UnitUid   uint          `json:"unitUid,omitempty"`
	IsOffline bool          `json:"isOffline,omitempty"`
	IsReady   bool          `json:"isReady,omitempty"`
}

func (p PlayerInfo) Clone() *PlayerInfo {
	return &p
}
