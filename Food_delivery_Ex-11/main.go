package main

import (
	"fmt"
	"food_delivery_service/component"
	"food_delivery_service/component/uploadprovider"
	"food_delivery_service/middleware"
	"food_delivery_service/modules/food/foodtransport/ginfood"
	"food_delivery_service/modules/upload/uploadtransport/ginupload"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := goDotEnvVariable("DBCONNECTIONSTR")
	s3BucketName := goDotEnvVariable("S3BUCKETNAME")
	s3Region := goDotEnvVariable("S3REGION")
	s3APIKey := goDotEnvVariable("S3APIKEY")
	s3SecretKey := goDotEnvVariable("S3SECRETKEY")
	s3Domain := goDotEnvVariable("S3DOMAIN")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("err1", err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln("err1", err)
	}
	fmt.Println(db)
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
