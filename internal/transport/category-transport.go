package transport

import (
	"mogshalla/internal/models"
	"mogshalla/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryTransport struct {
	service service.CategoryService
}

func NewCategoryTransport(service service.CategoryService) *CategoryTransport {
	return &CategoryTransport{service: service}
}

func (p *CategoryTransport) RegisterRoures(r *gin.Engine) {
	r.GET("/categories", p.GetCategories)
	r.POST("/categories", p.CreateCategory)
	r.DELETE("/categories/:id",p.DeleteCategory)
}

func (p *CategoryTransport) CreateCategory(c *gin.Context) {
	var req models.CreateCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	category, err := p.service.Create(req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, category)
}

func (p *CategoryTransport) GetCategories(c *gin.Context) {
	categories, err := p.service.Get()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, categories)
}

func (p *CategoryTransport) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "невалидный ID категории",
		})
		return
	}

	if err := p.service.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "КАТЕГОРИЯ УДАЛЕНА",
	})
}
