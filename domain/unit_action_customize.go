package domain

func (u *Unit) Customize(customization *UnitCustomization) *ActionResult {
	u.Customization = customization
	return NewActionResult().WithResult(ResultAccomplished)
}
