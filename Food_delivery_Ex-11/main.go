package main

import (
	"food_delivery_service/component"
	"food_delivery_service/component/uploadprovider"
	"food_delivery_service/middleware"
	"food_delivery_service/modules/food/foodtransport/ginfood"
	"food_delivery_service/modules/upload/uploadtransport/ginupload"
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

	s3BucketName := "food-go"
	s3Region := "ap-southeast-1"
	s3APIKey := "AKIAUS2KPLDJUOXE22X2"
	s3SecretKey := "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"
	s3Domain := "https://d1n632kj3y4onx.cloudfront.net"

	// s3BucketName := os.Getenv("S3BucketName")
	// s3Region := os.Getenv("S3Region")
	// s3APIKey := os.Getenv("S3APIKey")
	// s3SecretKey := os.Getenv("S3SecretKey")
	// s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(db)
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", ginupload.Upload(appCtx))
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
