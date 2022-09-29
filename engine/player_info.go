package engine

type PlayerInfo struct {
	Nickname  string        `json:"nickname"`
	Class     GameUnitClass `json:"class"`
	Level     uint          `json:"level"`
	UnitUid   uint          `json:"unitUid,omitempty"`
	IsOffline bool          `json:"isOffline,omitempty"`
	IsHost    bool          `json:"isHost,omitempty"`
}

func (p PlayerInfo) Clone() *PlayerInfo {
	return &p
}
