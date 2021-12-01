package account

type SocialEntity struct {
	Type     string    `gorm:"column:type"`
	Account     string    `gorm:"column:account"`
}
