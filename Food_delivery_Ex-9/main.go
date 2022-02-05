package main

import (
	"fmt"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodtransport/ginfood"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Food struct {
	Id           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	RestaurantId int    `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Status       int    `json:"status" gorm:"column:status;"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodUpdate struct {
	Name         *string `json:"name" gorm:"column:name;"`
	RestaurantId *int    `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Status       *int    `json:"status" gorm:"column:status;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456789@tcp(127.0.0.1:3366)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

	//dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	foods := r.Group("/foods")
	{
		// create food
		foods.POST("", ginfood.CreateFood(appCtx))

		// list foods
		foods.GET("", ginfood.ListFood(appCtx))

		//Get food by id
		foods.GET("/:id", ginfood.GetFood(appCtx))

		//Update food by id
		foods.PATCH("/:id", ginfood.UpdateFood(appCtx))

		//Delete food by id
		foods.DELETE("/:id", ginfood.DeleteFood(appCtx))
	}

	return r.Run()
}
