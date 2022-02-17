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

func CreateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.GenUID(common.DbTypeCategory)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
