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

func ListFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter foodmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewListFoodBiz(store)

		result, err := biz.ListFood(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
