package domain

type UnitProgress struct {
	Level                uint `json:"level"`
	Experience           uint `json:"experience"`
	ExperienceNext       uint `json:"experienceNext,omitempty"`
	AttributesPoints     uint `json:"attributesPoints,omitempty"`
	BaseAttributesPoints uint `json:"baseAttributesPoints,omitempty"`
}
