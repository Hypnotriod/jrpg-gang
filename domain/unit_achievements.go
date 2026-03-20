package domain

type UnitAchievement string

type UnitAchievements map[UnitAchievement]int

func (a UnitAchievements) Set(achievements UnitAchievements) {
	for achievement, value := range achievements {
		a[achievement] = value
	}
}

func (a UnitAchievements) Merge(achievements UnitAchievements) {
	for achievement, value := range achievements {
		a[achievement] += value
	}
}

func (a UnitAchievements) Test(achievements UnitAchievements) bool {
	for achievement, value := range achievements {
		if value < 0 {
			if a[achievement] >= value*-1 {
				return false
			}
		} else if a[achievement] < value {
			return false
		}
	}
	return true
}
