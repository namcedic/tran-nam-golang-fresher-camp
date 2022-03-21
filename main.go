package main

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/component/uploadprovider"
	"food_delivery_service/memcache"
	"food_delivery_service/middleware"
	"food_delivery_service/modules/category/categorytransport/gincategory"
	"food_delivery_service/modules/food/foodtransport/ginfood"
	"food_delivery_service/modules/restaurant/restauranttransport/ginrestaurant"
	"food_delivery_service/modules/restaurantlike/transport/ginrestaurantlike"
	"food_delivery_service/modules/upload/uploadtransport/ginupload"
	"food_delivery_service/modules/user/userstorage"
	"food_delivery_service/modules/user/usertransport/ginuser"
	"food_delivery_service/modules/userlike/transport/ginuserlike"
	"food_delivery_service/pubsub/pblocal"
	"food_delivery_service/skio"
	"food_delivery_service/subscriber"
	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("md.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	dsn := goDotEnvVariable("DBCONNECTIONSTR")

	s3BucketName := goDotEnvVariable("S3BUCKETNAME")
	s3Region := goDotEnvVariable("S3REGION")
	s3APIKey := goDotEnvVariable("S3APIKEY")
	s3SecretKey := goDotEnvVariable("S3SECRETKEY")
	s3Domain := goDotEnvVariable("S3DOMAIN")
	secretKey := goDotEnvVariable("LOGINSECRETKEY")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub())

	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	r := gin.Default()

	//subscriber.Setup(appCtx)
	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}
	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatalln(err)
	}

	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "true",
		})
	})

	// CRUD
	r.StaticFile("/demo/", "./demo.html")

	v1 := r.Group("/v1")
	v1.POST("/upload", middleware.RequiredAuth(appCtx, userCachingStore), ginupload.Upload(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx, userCachingStore), ginuser.GetProfile(appCtx))
	v1.GET("/users/:id/liked-restaurants", middleware.RequiredAuth(appCtx, userCachingStore), ginuserlike.ListRestaurant(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx, userCachingStore))
	{
		// create Restaurant
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		// list restaurants
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		//Get restaurant by id
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		//Update Restaurant by id
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		//Delete Restaurant by id
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		// Get Restaurant likes
		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))
	}

	foods := v1.Group("/foods", middleware.RequiredAuth(appCtx, userCachingStore))
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

	categories := v1.Group("/categories", middleware.RequiredAuth(appCtx, userCachingStore))
	{
		// create food
		categories.POST("", gincategory.CreateCategory(appCtx))

		// list foods
		categories.GET("", gincategory.ListCategory(appCtx))

		//Get food by id
		categories.GET("/:id", gincategory.GetCategory(appCtx))

		//Update food by id
		categories.PATCH("/:id", gincategory.UpdateCategory(appCtx))

		//Delete food by id
		categories.DELETE("/:id", gincategory.DeleteCategory(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	je, err := jg.NewExporter(jg.Options{
		AgentEndpoint: "localhost:6831",
		Process:       jg.Process{ServiceName: "Food-Delivery"},
	})

	if err != nil {
		log.Println(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})

	return http.ListenAndServe(
		":8080",
		&ochttp.Handler{
			Handler: r,
		},
	)
	//return r.Run()
}
