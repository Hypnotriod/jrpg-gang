package engine

type PlayerInfo struct {
	Id        UserId        `json:"-"`
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
