package transport

import (

	"mogshalla/internal/models"
	"mogshalla/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductTransport struct {
	service service.ProductService
}

func NewProductTransport(service service.ProductService) *ProductTransport {
	return &ProductTransport{service: service}
}

func (p *ProductTransport) RegisterRouters(r *gin.Engine) {
	r.GET("/categories/:id/products", p.GetByCategoriesID)
	r.POST("/products", p.CreateProduct)
	r.PATCH("/products/:id", p.UpdateProduct)
	r.DELETE("/products/:id", p.DeleteProduct)
}

func (p *ProductTransport) CreateProduct(c *gin.Context) {
	var req models.CreateProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	product, err := p.service.Create(&req)
	
	if err != nil {
		c.JSON(401, err.Error())
		return
	}


	c.JSON(200, product)
}

func (p *ProductTransport) GetByCategoriesID(c *gin.Context) {
	idSrt := c.Param("id")
	id, err := strconv.ParseUint(idSrt, 10, 64)

	if err != nil {
		c.JSON(400, err.Error())
	}

	products, err := p.service.Get(uint(id))

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, products)
}

func (p *ProductTransport) UpdateProduct(c *gin.Context) {
	idSrt := c.Param("id")
	id, err := strconv.ParseUint(idSrt, 10, 64)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.UpdateProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := p.service.Update(uint(id), &req); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "ПРОДУКТ ОБНОВЛЕН",
	})
}

func (p *ProductTransport) DeleteProduct(c *gin.Context) {
	idSrt := c.Param("id")
	id, err := strconv.ParseUint(idSrt, 10, 64)	
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := p.service.Delete(uint(id)); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "ПРОДУКТ УДАЛЕН",
	})
}
