package ginfood

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewGetFoodBiz(store)

		data, err := biz.GetFood(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
