package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id          int    `json:"id,omitempty" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Address     string `json:"address" gorm:"column:address;"`
	PhoneNumber int    `json:"phone_number" gorm:"column:phone_number;"`
}

type RestaurantUpdate struct {
	Id          *int    `json:"id,omitempty" gorm:"column:id;"`
	Name        *string `json:"name" gorm:"column:name;"`
	Address     *string `json:"address" gorm:"column:address;"`
	PhoneNumber *int    `json:"phone_number" gorm:"column:phone_number;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456789@tcp(127.0.0.1:3366)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

	//dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Insert new restaurant
	newRestaurant := Restaurant{Name: "New", Address: "123 Phố Huế", PhoneNumber: 383838}

	if err := db.Create(&newRestaurant); err != nil {
		fmt.Println(err)
	}

	// find all resaults
	var restaurants []Restaurant
	db.Where("phone_number=? ", 383838).Find(&restaurants)
	db.Where("address=? ", "123 Phố Huế").Find(&restaurants)
	fmt.Println(restaurants)

	// find first line
	var restaurant Restaurant
	if err := db.Where("id=? ", 1).First(&restaurant); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(restaurant)
	}

	// Delete
	db.Table(Restaurant{}.TableName()).Where("id = 2").Delete(nil)
	fmt.Println(restaurants)

	// Update
	restaurant.Address = "126 Lò Đúc"
	db.Table(Restaurant{}.TableName()).Where("id = 1").Updates(&restaurant)
	fmt.Println(restaurants)

	// Update with nil value
	newAddress := ""
	db.Table(Restaurant{}.TableName()).Where("id = 3").Updates(&RestaurantUpdate{Address: &newAddress})
	fmt.Println(restaurants)
}
