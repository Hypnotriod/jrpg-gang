package domain

type Damage struct {
	Stabbing float32 `json:"stabbing,omitempty"`
	Cutting  float32 `json:"cutting,omitempty"`
	Crushing float32 `json:"crushing,omitempty"`
	Fire     float32 `json:"fire,omitempty"`
	Cold     float32 `json:"cold,omitempty"`
	Lighting float32 `json:"lighting,omitempty"`
	Fear     float32 `json:"fear,omitempty"`
	Poison   float32 `json:"poison,omitempty"`
	Curse    float32 `json:"curse,omitempty"`
	Stunning float32 `json:"stunning,omitempty"`
	// todo
}
