package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/turamant/go-worehouse/models"
)

type ArrivalController struct {
	DB *models.DB
}

func NewArrivalController(db *models.DB) *ArrivalController {
	return &ArrivalController{DB: db}
}

func (ac *ArrivalController) CreateArrival(c echo.Context) error {
	// Логика создания нового прибытия
	return c.JSON(http.StatusCreated, nil)
}

func (ac *ArrivalController) UpdateArrival(c echo.Context) error {
	// Логика обновления прибытия
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (ac *ArrivalController) DeleteArrival(c echo.Context) error {
	// Логика удаления прибытия
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (ac *ArrivalController) GetArrivals(c echo.Context) error {
	// Логика получения списка прибытий
	return c.JSON(http.StatusOK, []string{"Arrival 1", "Arrival 2", "Arrival 3"})
}