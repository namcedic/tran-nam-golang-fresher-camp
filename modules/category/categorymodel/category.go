package categorymodel

import (
	"food_delivery_service/common"
	"strings"
)

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Status          int           `json:"status" gorm:"column:status;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryUpdate struct {
	Name        *string       `json:"name" gorm:"column:name;"`
	Description *string       `json:"description" gorm:"column:description;"`
	Status      *int          `json:"status" gorm:"column:status;"`
	Icon        *common.Image `json:"icon" gorm:"column:icon;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Status          int           `json:"status" gorm:"column:status;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (cat *CategoryCreate) Validate() error {
	cat.Name = strings.TrimSpace(cat.Name)
	if len(cat.Name) == 0 {
		return ErrNameCannotBeEmpty
	}
	return nil
}

func (data *Category) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCategory)
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "category name can't be blank", "ErrNameCannotBeEmpty")
)
