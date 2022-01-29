package foodmodel

type Food struct {
	Id     int    `json:"id,omitempty" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Status int    `json:"status" gorm:"column:status"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodUpdate struct {
	Name   *string `json:"name" gorm:"column:name"`
	Status *int    `json:"status" gorm:"column:status"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

type FoodCreate struct {
	Name   string `json:"name" gorm:"column:name"`
	Status int    `json:"status" gorm:"column:status"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}
