package transport

import (
	"mogshalla/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(
	r *gin.Engine,
	categoryServ service.CategoryService,
	productServ service.ProductService,
) {
	categoryHandler := NewCategoryTransport(categoryServ)
	categoryHandler.RegisterRoures(r)

	productHandler := NewProductTransport(productServ)
	productHandler.RegisterRouters(r)
}
