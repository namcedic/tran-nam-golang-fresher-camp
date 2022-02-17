package common

type SimpleRestaurant struct {
	SQLModel   `json:",inline"`
	Name       string `json:"name" gorm:"column:name;"`
	Addr       string `json:"address" gorm:"column:addr;"`
	Logo       *Image `json:"logo" gorm:"column:logo;"`
	LikedCount int    `json:"liked_count" gorm:"-"`
}

func (SimpleRestaurant) TableName() string {
	return "restaurants"
}

func (u *SimpleRestaurant) Mask(isAdmin bool) {
	u.GenUID(DbTypeRestaurant)
}
