package ginfood

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/food/foodbiz"
	"food_delivery_service/modules/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewDeleteFoodBiz(store)

		if err := biz.DeleteFood(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
