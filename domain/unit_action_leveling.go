package domain

import "math"

func (u *Unit) LevelUp() *ActionResult {
	result := NewActionResult()
	if u.Stats.Progress.Experience < u.Stats.Progress.ExperienceNext {
		return result.WithResult(ResultNotEnoughResources)
	}
	u.Stats.Progress.Level++
	u.Stats.Progress.Experience -= u.Stats.Progress.ExperienceNext
	u.Stats.Progress.AttributesPoints += LEVEL_UP_ATTRIBUTES_POINTS
	u.Stats.Progress.BaseAttributesPoints += LEVEL_UP_BASE_ATTRIBUTES_POINTS
	u.Stats.Progress.ExperienceNext = u.NextLevelExp()
	return result.WithResult(ResultAccomplished)
}

func (u *Unit) NextLevelExp() uint {
	raw := EXPERIENCE_SEED * math.Pow(float64(u.Stats.Progress.Level), EXPERIENCE_POWER)
	return uint(raw/EXPERIENCE_SCALE) * uint(EXPERIENCE_SCALE)
}

func (u *Unit) SkillUp(skill ActionProperty) *ActionResult {
	result := NewActionResult()
	switch skill {
	case AtionPropertyStrength:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Strength) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyPhysique:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Physique) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyAgility:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Agility) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyEndurance:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Endurance) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyIntelligence:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Intelligence) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyInitiative:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Initiative) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyLuck:
		if !u.tryToIncreaseAttribute(&u.Stats.Attributes.Luck) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyHealth:
		if !u.tryToIncreaseBaseAttribute(&u.Stats.BaseAttributes.Health) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyStamina:
		if !u.tryToIncreaseBaseAttribute(&u.Stats.BaseAttributes.Stamina) {
			return result.WithResult(ResultNotEnoughResources)
		}
	case AtionPropertyMana:
		if !u.tryToIncreaseBaseAttribute(&u.Stats.BaseAttributes.Mana) {
			return result.WithResult(ResultNotEnoughResources)
		}
	default:
		return result.WithResult(ResultNotFound)
	}
	return result.WithResult(ResultAccomplished)
}

func (u *Unit) tryToIncreaseAttribute(attribute *float32) bool {
	price := uint(*attribute)/ATTRIBUTE_COST_STEP + 1
	if u.Stats.Progress.AttributesPoints < price {
		return false
	}
	u.Stats.Progress.AttributesPoints -= price
	*attribute++
	return true
}

func (u *Unit) tryToIncreaseBaseAttribute(attribute *float32) bool {
	price := uint(*attribute)/BASE_ATTRIBUTE_COST_STEP + 1
	if u.Stats.Progress.BaseAttributesPoints < price {
		return false
	}
	u.Stats.Progress.BaseAttributesPoints -= price
	*attribute++
	return true
}
