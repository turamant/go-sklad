// main.go
package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/turamant/go-worehouse/models"
	"github.com/turamant/go-worehouse/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:password@tcp(172.17.0.2:3306)/"
	db, err := models.NewDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверка подключения
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connected to database")

	// Создание экземпляра Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Настройка маршрутов
	routes.SetupRoutes(e, db)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8000"))
}

