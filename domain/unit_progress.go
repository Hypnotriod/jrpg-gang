package domain

type UnitProgress struct {
	Level                uint `json:"level" bson:"level"`
	Experience           uint `json:"experience" bson:"experience"`
	ExperienceNext       uint `json:"experienceNext,omitzero" bson:"-"`
	AttributesPoints     uint `json:"attributesPoints,omitzero" bson:"attributes_points,omitempty"`
	BaseAttributesPoints uint `json:"baseAttributesPoints,omitzero" bson:"baseAttributes_points,omitempty"`
}
