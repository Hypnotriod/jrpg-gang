package engine

import "fmt"

type Spot struct {
	Name        string      `json:"name"`
	Battlefield Battlefield `json:"battlefield"`
}

func (s Spot) String() string {
	return fmt.Sprintf(
		"%s, battlefield: {%v}",
		s.Name,
		s.Battlefield,
	)
}
