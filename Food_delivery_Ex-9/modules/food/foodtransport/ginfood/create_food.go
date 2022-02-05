package ginfood

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodmodel"
	"food_delivery_service/modules/food/foodstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodmodel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewCreateFoodBiz(store)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
