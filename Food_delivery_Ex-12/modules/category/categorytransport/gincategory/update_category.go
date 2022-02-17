package gincategory

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/category/categorybiz"
	"food_delivery_service/modules/category/categorymodel"
	"food_delivery_service/modules/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data categorymodel.CategoryUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err := biz.UpdateCategory(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
