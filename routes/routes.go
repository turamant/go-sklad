// routes/routes.go
package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/turamant/go-worehouse/controllers"
	"github.com/turamant/go-worehouse/models"
)

func SetupRoutes(e *echo.Echo, db *models.DB) *echo.Echo {

	productController := controllers.NewProductController(db)
	arrivalController := controllers.NewArrivalController(db)

	e.POST("/products", productController.CreateProduct)
	e.PUT("/products/:id", productController.UpdateProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)
	e.GET("/products", productController.GetProducts)

	e.POST("/arrivals", arrivalController.CreateArrival)
	e.PUT("/arrivals/:id", arrivalController.UpdateArrival)
	e.DELETE("/arrivals/:id", arrivalController.DeleteArrival)
	e.GET("/arrivals", arrivalController.GetArrivals)

	return e
}
