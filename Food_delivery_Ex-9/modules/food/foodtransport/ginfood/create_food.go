package ginfood

import (
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodmodel"
	"food_delivery_service/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFood(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewCreateFoodBiz(store)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
