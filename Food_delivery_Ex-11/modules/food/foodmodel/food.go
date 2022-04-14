package foodmodel

import (
	"food_delivery_service/common"
	"strings"
)

const EntityName = "Food"

type Food struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	RestaurantId    int            `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Status          int            `json:"status" gorm:"column:status;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	DetailImg       *common.Images `json:"detail_image" gorm:"column:detail_image;"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodUpdate struct {
	Name         *string        `json:"name" gorm:"column:name;"`
	RestaurantId *int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Status       *int           `json:"status" gorm:"column:status;"`
	Logo         *common.Image  `json:"logo" gorm:"column:logo;"`
	DetailImg    *common.Images `json:"detail_image" gorm:"column:detail_image;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

type FoodCreate struct {
	Id           int            `json:"id" gorm:"column:id;"`
	Name         string         `json:"name" gorm:"column:name;"`
	RestaurantId int            `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Status       int            `json:"status" gorm:"column:status;"`
	Logo         *common.Image  `json:"logo" gorm:"column:logo;"`
	DetailImg    *common.Images `json:"detail_image" gorm:"column:detail_image;"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}

func (res *FoodCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "food name can't be blank", "ErrNameCannotBeEmpty")
)
