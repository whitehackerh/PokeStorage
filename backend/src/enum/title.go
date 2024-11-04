package enum

const (
	TitleSwordShield   = 1
	TitleScarletViolet = 2
)

func IsSwordShield(titleId int) bool {
	return titleId == TitleSwordShield
}

func IsScarletViolet(titleId int) bool {
	return titleId == TitleScarletViolet
}
