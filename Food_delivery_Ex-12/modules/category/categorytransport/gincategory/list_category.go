package gincategory

import (
	"food_delivery_service/common"
	"food_delivery_service/component"
	"food_delivery_service/modules/category/categorybiz"
	"food_delivery_service/modules/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewListCategoryBiz(store)

		result, err := biz.ListCategory(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
