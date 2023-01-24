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
