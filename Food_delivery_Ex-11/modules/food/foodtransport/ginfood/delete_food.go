package ginfood

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewDeleteFoodBiz(store)

		if err := biz.DeleteFood(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
