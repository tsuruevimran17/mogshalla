package main

import (
	"mogshalla/internal/config"
	"mogshalla/internal/models"
	"mogshalla/internal/repository"
	"mogshalla/internal/service"
	"mogshalla/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabase()

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Product{},
	); err != nil {
		panic(err)
	}

	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)

	categoryServ := service.NewCategoryService(categoryRepo)
	productServ := service.NewProductService(productRepo)

	router := gin.Default()

	transport.RegisterRouters(
		router,
		categoryServ,
		productServ,
	)

	router.Run(":8080")
}
