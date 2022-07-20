package engine

type PlayerInfo struct {
	Nickname  string        `json:"nickname"`
	Class     GameUnitClass `json:"class"`
	Level     uint          `json:"level"`
	IsOffline bool          `json:"isOffline,omitempty"`
}
