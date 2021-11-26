package domain

import "fmt"

type UnitProgress struct {
	Level      float32 `json:"level"`
	Experience float32 `json:"experience"`
}

func (p UnitProgress) String() string {
	return fmt.Sprintf("level: %g, experience: %g",
		p.Level, p.Experience)
}
