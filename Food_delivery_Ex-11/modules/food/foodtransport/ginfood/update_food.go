package ginfood

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodmodel"
	"food_delivery_service/modules/food/foodstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data foodmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewUpdateFoodBiz(store)

		if err := biz.UpdateFood(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
