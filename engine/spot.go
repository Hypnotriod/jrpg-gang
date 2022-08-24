package engine

import "fmt"

type Spot struct {
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Battlefield Battlefield `json:"battlefield"`
}

func (s Spot) String() string {
	return fmt.Sprintf(
		"%s, code: %s, battlefield: {%v}",
		s.Name,
		s.Code,
		s.Battlefield,
	)
}
