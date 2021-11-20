package domain

type Unit struct {
	State       UnitState               `json:"state"`
	Stats       UnitStats               `json:"stats"`
	Damage      []DamageImpact          `json:"damage"`
	Enhancement []UnitEnhancementImpact `json:"enhancement"`
	Items       []Item                  `json:"items"`
}

type UnitStats struct {
	Progress       UnitProgress       `json:"progress"`
	BaseAttributes UnitBaseAttributes `json:"baseAttributes"`
	Attributes     UnitAttributes     `json:"attributes"`
	Resistance     UnitResistance     `json:"resistance"`
}

type UnitBaseAttributes struct {
	Health  float32 `json:"health"`
	Stamina float32 `json:"stamina"`
	Mana    float32 `json:"mana"`
}

type UnitState struct {
	UnitBaseAttributes
}

type UnitProgress struct {
	Level      float32 `json:"level"`
	Experience float32 `json:"experience"`
}

type UnitResistance struct {
	Damage
}

type UnitAttributes struct {
	Strength     float32 `json:"strength"`
	Physique     float32 `json:"physique"`
	Dexterity    float32 `json:"dexterity"`
	Endurance    float32 `json:"endurance"`
	Intelligence float32 `json:"intelligence"`
	Holy         float32 `json:"holy"`
	Luck         float32 `json:"luck"`
}

type UnitEnhancement struct {
	UnitBaseAttributes `json:"unitBaseAttributes"`
	UnitAttributes     `json:"unitAttributes"`
	UnitResistance     `json:"unitResistance"`
}
