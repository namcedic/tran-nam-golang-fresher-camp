package main

import (
	"errors"
	"fmt"
	"food_delivery_service/modules/food/foodtransport/ginfood"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	foods := r.Group("/foods")
	{
		foods.POST("", ginfood.CreateFood(db))

		// list foods
		foods.GET("", func(c *gin.Context) {
			var listFood []Food
			type Filter struct {
				Status int `json:"status" form:"status"`
			}
			var filter Filter
			if err := c.ShouldBind(&filter); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			newDB := db
			if filter.Status > 0 {
				newDB = db.Where("status = ?", filter.Status)
			}
			if err := newDB.Find(&listFood).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, listFood)
		})

		//Get food by id
		foods.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data Food
			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		//Update food by id
		foods.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			var data FoodUpdate
			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": errors.New("update failed"),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})
		//Delete food by id

		foods.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			if err := db.Table(Food{}.TableName()).
				Where("id = ?", id).Delete(nil).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"ok": "1"})
		})
	}

	return r.Run()
}

func CreateFood() {
	panic("unimplemented")
}
