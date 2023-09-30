package domain

type UnitAchievement string

type UnitAchievements map[UnitAchievement]uint

func (a UnitAchievements) Accumulate(achievements []UnitAchievement) {
	for _, achievement := range achievements {
		a[achievement] += 1
	}
}

func (a UnitAchievements) Merge(achievements UnitAchievements) {
	for achievement, value := range achievements {
		a[achievement] += value
	}
}

func (a UnitAchievements) Test(achievements UnitAchievements) bool {
	for achievement, value := range achievements {
		if a[achievement] < value {
			return false
		}
	}
	return true
}
