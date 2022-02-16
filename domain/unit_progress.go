package domain

import "fmt"

type UnitProgress struct {
	Level      uint `json:"level"`
	Experience uint `json:"experience"`
}

func (p UnitProgress) String() string {
	return fmt.Sprintf("level: %d, experience: %d",
		p.Level, p.Experience)
}
