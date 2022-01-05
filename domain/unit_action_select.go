package domain

func (u *Unit) SelectAmmunition(uid uint) {
	u.Inventory.SelectAmmunition(uid)
}
