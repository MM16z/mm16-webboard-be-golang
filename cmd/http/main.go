package main

import (
	"log"
	"time"

	"github.com/MM16z/mm16-webboard-be-golang/internal/adapter/handler/http"
	"github.com/MM16z/mm16-webboard-be-golang/internal/adapter/storage/postgres"
	"github.com/MM16z/mm16-webboard-be-golang/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	db := postgres.InitDB()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	initApp := struct {
		Message string
	}{
		Message: "hello world! " + util.FormatTimeThaiEng(time.Now()),
	}

	http.InitRouter(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(initApp)
	})

	for _, route := range app.GetRoutes() {
		log.Printf("Method: %s, Path: %s\n", route.Method, route.Path)
	}

	log.Println("Starting server on http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
