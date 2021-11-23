package domain

func AsEquipment(item interface{}) (*Equipment, bool) {
	var r *Equipment
	var ok bool = false
	switch v := item.(type) {
	case Equipment:
		r, ok = &v, true
	case Weapon:
		r, ok = &v.Equipment, true
	case Armor:
		r, ok = &v.Equipment, true
	}
	return r, ok
}
