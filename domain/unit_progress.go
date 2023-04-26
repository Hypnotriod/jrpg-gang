package domain

type UnitProgress struct {
	Level                uint `json:"level" bson:"level"`
	Experience           uint `json:"experience" bson:"experience"`
	ExperienceNext       uint `json:"experienceNext,omitempty" bson:"-"`
	AttributesPoints     uint `json:"attributesPoints,omitempty" bson:"attributes_points,omitempty"`
	BaseAttributesPoints uint `json:"baseAttributesPoints,omitempty" bson:"baseAttributes_points,omitempty"`
}
