package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/turamant/go-worehouse/models"
)

type ProductController struct {
	DB *models.DB
}

func NewProductController(db *models.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	// Логика создания нового продукта
	return c.JSON(http.StatusCreated, nil)
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	// Логика обновления продукта
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	// Логика удаления продукта
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (pc *ProductController) GetProducts(c echo.Context) error {
	// Логика получения списка продуктов
	return c.JSON(http.StatusOK, []string{"Product 1", "Product 2", "Product 3"})
}