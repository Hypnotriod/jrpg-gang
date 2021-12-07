package domain

import "fmt"

type Disposable struct {
	Item
	Impact      []DamageImpact          `json:"impact,omitempty"`
	Enhancement []UnitEnhancementImpact `json:"enhancement,omitempty"`
}

func (d Disposable) String() string {
	return fmt.Sprintf(
		"%s, description: %s, enhancement: %v, impact: %v",
		d.Name,
		d.Description,
		d.Enhancement,
		d.Impact,
	)
}
