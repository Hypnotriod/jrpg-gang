package engine

import "fmt"

type Spot struct {
	Battlefield Battlefield `json:"battlefield"`
}

func (s Spot) String() string {
	return fmt.Sprintf(
		"battlefield: {%v}",
		s.Battlefield,
	)
}
